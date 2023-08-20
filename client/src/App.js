import React from 'react';
import { ApolloProvider } from '@apollo/client';
import client from './apolloClient';
import TodoList from './TodoList';
import CreateTodo from './CreateTodo';

function App() {
  return (
    <ApolloProvider client={client}>
      <div className="App">
        <h1>My Todos</h1>
        <CreateTodo />
        <TodoList />
      </div>
    </ApolloProvider>
  );
}

export default App;

