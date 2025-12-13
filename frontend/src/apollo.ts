import { ApolloClient, InMemoryCache, HttpLink } from '@apollo/client/core'

// HTTP connection to the API
// Use relative URL to go through nginx proxy which adds auth headers
const httpLink = new HttpLink({
  uri: '/query',
})

// Cache implementation
const cache = new InMemoryCache()

// Create the apollo client
export const apolloClient = new ApolloClient({
  link: httpLink,
  cache,
})
