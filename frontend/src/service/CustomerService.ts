import { apolloClient } from '@/apollo'
import gql from 'graphql-tag'

const GET_CUSTOMERS = gql`
  query GetCustomers {
    customers {
      id
      name
      country {
        name
        code
      }
      company
      date
      status
      verified
      activity
      representative {
        name
        image
      }
      balance
    }
  }
`

const GET_CUSTOMERS_LARGE = gql`
  query GetCustomersLarge {
    customersLarge {
      id
      name
      country {
        name
        code
      }
      company
      date
      status
      verified
      activity
      representative {
        name
        image
      }
      balance
    }
  }
`

const GET_CUSTOMERS_MEDIUM = gql`
  query GetCustomersMedium {
    customersMedium {
      id
      name
      country {
        name
        code
      }
      company
      date
      status
      verified
      activity
      representative {
        name
        image
      }
      balance
    }
  }
`

export const CustomerService = {
  async getCustomersSmall() {
    const { data } = await apolloClient.query({
      query: GET_CUSTOMERS
    })
    return data.customers
  },

  async getCustomersMedium() {
    const { data } = await apolloClient.query({
      query: GET_CUSTOMERS_MEDIUM
    })
    return data.customersMedium
  },

  async getCustomersLarge() {
    const { data } = await apolloClient.query({
      query: GET_CUSTOMERS_LARGE
    })
    return data.customersLarge
  }
}
