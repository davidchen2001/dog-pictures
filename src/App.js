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

export default App;
