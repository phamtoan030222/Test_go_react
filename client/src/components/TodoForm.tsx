import {
  Button,
  Flex,
  Input,
  Spinner,
  useDisclosure,
  Modal,
  ModalOverlay,
  ModalContent,
  ModalHeader,
  ModalFooter,
  ModalBody,
  ModalCloseButton,
  FormControl,
  FormLabel,
} from "@chakra-ui/react";
import { useMutation, useQueryClient } from "@tanstack/react-query";
import { useState } from "react";
import { IoMdAdd } from "react-icons/io";
import { BASE_URL } from "../App";

const TodoForm = () => {
  const queryClient = useQueryClient();
  const { isOpen, onOpen, onClose } = useDisclosure();

  const [title, setTitle] = useState("");
  const [description, setDescription] = useState("");

  const { mutate: createTodo, isPending: isCreating } = useMutation({
    mutationKey: ["createTodo"],
    mutationFn: async () => {
      const res = await fetch(BASE_URL + `/tasks`, {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ title, description }),
      });
      const data = await res.json();

      if (!res.ok) throw new Error(data.error || "Something went wrong");

      return data;
    },
    onSuccess: () => {
      queryClient.invalidateQueries({ queryKey: ["todos"] });
      setTitle("");
      setDescription("");
      onClose(); // đóng modal khi thêm thành công
    },
    onError: (error: any) => {
      alert(error.message);
    },
  });

  return (
    <>
      {/* Nút mở modal */}
      <Flex my={3}>
        <Button
          leftIcon={<IoMdAdd size={20} />}
          colorScheme="blue"
          onClick={onOpen}
        >
          Thêm Task
        </Button>
      </Flex>

      {/* Modal thêm task */}
      <Modal isOpen={isOpen} onClose={onClose} isCentered>
        <ModalOverlay />
        <ModalContent>
          <ModalHeader>Thêm công việc mới</ModalHeader>
          <ModalCloseButton />
          <ModalBody>
            <FormControl mb={3}>
              <FormLabel>Tiêu đề</FormLabel>
              <Input
                value={title}
                onChange={(e) => setTitle(e.target.value)}
                placeholder="Nhập tiêu đề..."
              />
            </FormControl>

            <FormControl>
              <FormLabel>Mô tả</FormLabel>
              <Input
                value={description}
                onChange={(e) => setDescription(e.target.value)}
                placeholder="Nhập mô tả..."
              />
            </FormControl>
          </ModalBody>

          <ModalFooter>
            <Button variant="ghost" mr={3} onClick={onClose}>
              Hủy
            </Button>
            <Button
              colorScheme="blue"
              onClick={() => createTodo()}
              isDisabled={!title.trim()}
            >
              {isCreating ? <Spinner size="sm" /> : "Lưu"}
            </Button>
          </ModalFooter>
        </ModalContent>
      </Modal>
    </>
  );
};

export default TodoForm;
