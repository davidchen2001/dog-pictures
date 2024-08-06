import { useState, useEffect } from "react";
import axios from "axios";
import {
  Card,
  CardHeader,
  CardBody,
  CardFooter,
  Image,
  Heading,
  Text,
  Stack,
  Divider,
  ButtonGroup,
  Button,
} from "@chakra-ui/react";

import { BASE_URL } from "../../App";

function DogCard() {
  useEffect(() => {
    const getData = async () => {
      const query = await axios
        .get(BASE_URL + "/dog")
        .then(function (response) {
          console.log(response);
        });
    };
    getData().then((data) => {
      console.log(data);
    });
  }, []);

  return (
    <Card maxW="sm">
      <CardBody>
        <Image
          src="https://images.dog.ceo/breeds/spaniel-welsh/n02102177_2755.jpg"
          alt="Green double couch with wooden legs"
          borderRadius="lg"
        />
        <Stack mt="6" spacing="3">
          <Heading size="md">Dog</Heading>
        </Stack>
      </CardBody>
      <CardFooter>
        <ButtonGroup spacing="2">
          <Button variant="solid" colorScheme="blue">
            Next Dog
          </Button>
        </ButtonGroup>
      </CardFooter>
    </Card>
  );
}

export default DogCard;
