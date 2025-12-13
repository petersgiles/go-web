import { apolloClient } from '@/apollo'
import gql from 'graphql-tag'

const GET_PHOTOS = gql`
  query GetPhotos {
    photos {
      title
      itemImageSrc
      thumbnailImageSrc
      alt
    }
  }
`

export const PhotoService = {
  async getImages() {
    const { data } = await apolloClient.query({
      query: GET_PHOTOS
    })
    return data.photos
  }
}
