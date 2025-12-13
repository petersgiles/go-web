import { apolloClient } from '@/apollo'
import gql from 'graphql-tag'

const GET_PRODUCTS = gql`
  query GetProducts {
    products {
      id
      code
      name
      description
      image
      price
      category
      quantity
      inventoryStatus
      rating
    }
  }
`

const GET_PRODUCTS_SMALL = gql`
  query GetProductsSmall {
    productsSmall {
      id
      code
      name
      description
      image
      price
      category
      quantity
      inventoryStatus
      rating
    }
  }
`

const GET_PRODUCTS_WITH_ORDERS = gql`
  query GetProductsWithOrders {
    productsWithOrders {
      id
      code
      name
      description
      image
      price
      category
      quantity
      inventoryStatus
      rating
      orders {
        id
        productCode
        date
        amount
        quantity
        customer
        status
      }
    }
  }
`

export const ProductService = {
  async getProductsSmall() {
    const { data } = await apolloClient.query({
      query: GET_PRODUCTS_SMALL
    })
    return data.productsSmall
  },

  async getProducts() {
    const { data } = await apolloClient.query({
      query: GET_PRODUCTS
    })
    return data.products
  },

  async getProductsMini() {
    return this.getProductsSmall()
  },

  async getProductsWithOrdersSmall() {
    const { data } = await apolloClient.query({
      query: GET_PRODUCTS_WITH_ORDERS
    })
    return data.productsWithOrders
  },

  async getProductsWithOrders() {
    const { data } = await apolloClient.query({
      query: GET_PRODUCTS_WITH_ORDERS
    })
    return data.productsWithOrders
  }
}
