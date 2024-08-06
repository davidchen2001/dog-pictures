import "./App.css";
import { ChakraProvider, Center } from "@chakra-ui/react";
import DogCard from "./components/DogCard/DogCard";

function App() {
  return (
    <ChakraProvider>
      <Center>
        <DogCard />
      </Center>
    </ChakraProvider>
  );
}

export const BASE_URL =
  process.env.NODE_ENV === "development" ? "http://localhost:6000/api" : "/api";

export default App;
