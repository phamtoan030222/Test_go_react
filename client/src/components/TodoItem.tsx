import {
  Badge,
  Box,
  Flex,
  Spinner,
  Text,
  useColorModeValue,
} from "@chakra-ui/react";
import { FaCheckCircle } from "react-icons/fa";
import { MdDelete } from "react-icons/md";
import type { Todo } from "./TodoList";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { BASE_URL } from "../App";

const TodoItem = ({ todo }: { todo: Todo }) => {
  const queryClient = useQueryClient();

  const { mutate: updateTodo, isPending: isUpdating } = useMutation({
    mutationKey: ["updateTodo"],
    mutationFn: async () => {
      if (todo.completed) return alert("Todo is already completed");
      const res = await fetch(BASE_URL + `/todos/${todo._id}`, {
        method: "PATCH",
      });
      const data = await res.json();
      if (!res.ok) throw new Error(data.error || "Something went wrong");
      return data;
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["todos"] }),
  });

  const { mutate: deleteTodo, isPending: isDeleting } = useMutation({
    mutationKey: ["deleteTodo"],
    mutationFn: async () => {
      const res = await fetch(BASE_URL + `/todos/${todo._id}`, {
        method: "DELETE",
      });
      const data = await res.json();
      if (!res.ok) throw new Error(data.error || "Something went wrong");
      return data;
    },
    onSuccess: () => queryClient.invalidateQueries({ queryKey: ["todos"] }),
  });

  // Dùng màu theo theme (light/dark mode)
  const completedColor = useColorModeValue("green.600", "green.200");
  const inProgressColor = useColorModeValue("yellow.600", "yellow.200");

  return (
    <Flex gap={2} alignItems="center">
      <Flex
        flex={1}
        alignItems="center"
        border="1px"
        borderColor={useColorModeValue("gray.300", "gray.600")}
        p={2}
        borderRadius="lg"
        justifyContent="space-between"
      >
        <Text
          color={todo.completed ? completedColor : inProgressColor}
          textDecoration={todo.completed ? "line-through" : "none"}
        >
          {todo.body}
        </Text>

        {todo.completed ? (
          <Badge ml={1} colorScheme="green">
            Done
          </Badge>
        ) : (
          <Badge ml={1} colorScheme="yellow">
            In Progress
          </Badge>
        )}
      </Flex>

      <Flex gap={2} alignItems="center">
        <Box color="green.400" cursor="pointer" onClick={() => updateTodo()}>
          {!isUpdating ? <FaCheckCircle size={20} /> : <Spinner size="sm" />}
        </Box>
        <Box color="red.400" cursor="pointer" onClick={() => deleteTodo()}>
          {!isDeleting ? <MdDelete size={25} /> : <Spinner size="sm" />}
        </Box>
      </Flex>
    </Flex>
  );
};

export default TodoItem;
