import {ApolloClient, InMemoryCache} from '@apollo/client';

const client = new ApolloClient({
    ui: 'http://localhost:8080/query',
    cache: new InMemoryCache()
});

export default client;
