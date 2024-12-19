import { Box, Container, Flex, Heading } from '@chakra-ui/react';
import Link from 'next/link';

export default function Header() {
  return (
    <Box px={4} bgColor='gray.100'>
      <Container maxW='container.lg'>
        <Flex
          as='header'
          py='4'
          justifyContent='space-between'
          alignItems='center'
        >
          <Link href='/' passHref>
            <Heading as='h1' fontSize='2xl' fontWeight='bold' cursor='pointer'>
              ポケポケ デッキ共有サイト
            </Heading>
          </Link>
        </Flex>
      </Container>
    </Box>
  );
}
