import { ref, computed } from 'vue'
import { useQuery } from '@vue/apollo-composable'
import gql from 'graphql-tag'

export interface User {
  id: string
  email: string
  name: string
  roles: string[]
  avatar?: string
  department?: string
  permissions: string[]
}

const CURRENT_USER_QUERY = gql`
  query CurrentUser {
    currentUser {
      id
      email
      name
      roles
      avatar
      department
      permissions
    }
  }
`

export function useUser() {
  const { result, loading, error, refetch } = useQuery(CURRENT_USER_QUERY, null, {
    fetchPolicy: 'cache-and-network',
  })

  const user = computed<User | null>(() => {
    return result.value?.currentUser || null
  })

  const isAuthenticated = computed(() => user.value !== null)

  const hasRole = (role: string) => {
    return user.value?.roles.includes(role) || false
  }

  const hasPermission = (permission: string) => {
    return user.value?.permissions.includes(permission) || false
  }

  const isAdmin = computed(() => hasRole('admin'))

  return {
    user,
    loading,
    error,
    isAuthenticated,
    hasRole,
    hasPermission,
    isAdmin,
    refetch,
  }
}
