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

const TodoItem = ({ task }: { task: Todo }) => {
  const queryClient = useQueryClient();

  // ---- Update task (optimistic) ----
  const { mutate: updateTask, isPending: isUpdating } = useMutation({
    mutationFn: async () => {
      const res = await fetch(`${BASE_URL}/tasks/${task.id}`, {
        method: "PATCH",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ completed: !task.completed }),
      });
      if (!res.ok) throw new Error("Failed to update");
      return res.json();
    },
    onMutate: async () => {
      await queryClient.cancelQueries({ queryKey: ["todos"] });

      const prevTasks = queryClient.getQueryData<Todo[]>(["tasks"]);

      queryClient.setQueryData<Todo[]>(["tasks"], (old) =>
        old
          ? old.map((t) =>
              t.id === task.id ? { ...t, completed: !t.completed } : t
            )
          : []
      );

      return { prevTasks };
    },
    onError: (_err, _vars, context) => {
      if (context?.prevTasks) {
        queryClient.setQueryData(["todos"], context.prevTasks);
      }
    },
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });

  // ---- Delete task (optimistic) ----
  const { mutate: deleteTask, isPending: isDeleting } = useMutation({
    mutationFn: async () => {
      const res = await fetch(`${BASE_URL}/tasks/${task.id}`, {
        method: "DELETE",
      });
      if (!res.ok) throw new Error("Failed to delete");
      return res.json();
    },
    onMutate: async () => {
      await queryClient.cancelQueries({ queryKey: ["todos"] });

      const prevTasks = queryClient.getQueryData<Todo[]>(["tasks"]);

      queryClient.setQueryData<Todo[]>(
        ["tasks"],
        (old) => old?.filter((t) => t.id !== task.id) || []
      );

      return { prevTasks };
    },
    onError: (_err, _vars, context) => {
      if (context?.prevTasks) {
        queryClient.setQueryData(["todos"], context.prevTasks);
      }
    },
    onSettled: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
    },
  });

  // ---- Colors ----
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
        <Box>
          <Text
            fontWeight="bold"
            color={task.completed ? completedColor : inProgressColor}
            textDecoration={task.completed ? "line-through" : "none"}
          >
            {task.title}
          </Text>
          <Text
            fontSize="sm"
            color={useColorModeValue("gray.600", "gray.400")}
            textDecoration={task.completed ? "line-through" : "none"}
          >
            {task.description}
          </Text>
        </Box>

        {task.completed ? (
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
        <Box color="green.400" cursor="pointer" onClick={() => updateTask()}>
          {!isUpdating ? <FaCheckCircle size={20} /> : <Spinner size="sm" />}
        </Box>
        <Box color="red.400" cursor="pointer" onClick={() => deleteTask()}>
          {!isDeleting ? <MdDelete size={25} /> : <Spinner size="sm" />}
        </Box>
      </Flex>
    </Flex>
  );
};

export default TodoItem;
