import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client/core'

// HTTP connection to the API
const httpLink = new HttpLink({
  // You should use an absolute URL here
  uri: 'http://localhost:8080/query',
})

// Cache implementation
const cache = new InMemoryCache()

// Create the apollo client
export const apolloClient = new ApolloClient({
  link: httpLink,
  cache,
})
