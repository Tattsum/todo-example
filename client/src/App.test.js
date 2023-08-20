import React from 'react';
import { ApolloProvider } from '@apollo/client';
import client from './apolloClient';
import TodoList from './TodoList';

function App() {
  return (
    <ApolloProvider client={client}>
      <TodoList />
    </ApolloProvider>
  );
}
