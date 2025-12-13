import { apolloClient } from '@/apollo'
import gql from 'graphql-tag'

const GET_COUNTRIES = gql`
  query GetCountries {
    countries {
      name
      code
    }
  }
`

export const CountryService = {
  async getCountries() {
    const { data } = await apolloClient.query({
      query: GET_COUNTRIES
    })
    return data.countries
  }
}
