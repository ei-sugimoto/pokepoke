import { newBackClient } from '@/utils/backClient';
import { Heading } from '@chakra-ui/react';

export default async function Home() {
  const client = newBackClient();
  const res = await fetch(client.URL + '/test');
  const msg = await res.json();
  console.log(msg);
  return (
    <Heading as={'h1'} fontSize='2xl'>
      デッキ一覧
    </Heading>
  );
}
