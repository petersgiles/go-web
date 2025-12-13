<script setup lang="ts">
import { useUser } from '@/composables/useUser'

const { user, loading } = useUser()
</script>

<template>
    <div class="layout-profile-menu hidden shadow-md">
        <div class="layout-profile-menu-content">
            <!-- Loading State -->
            <div v-if="loading" class="p-4 text-center">
                <i class="pi pi-spin pi-spinner text-2xl text-primary"></i>
                <p class="mt-2 text-sm text-muted-color">Loading profile...</p>
            </div>

            <!-- User Profile -->
            <div v-else-if="user" class="p-0">
                <!-- Header with Avatar -->
                <div class="bg-primary p-4 text-center">
                    <img 
                        v-if="user.avatar" 
                        :src="user.avatar" 
                        :alt="user.name"
                        class="w-20 h-20 rounded-full mx-auto mb-3 border-3 border-surface-0"
                    />
                    <div v-else class="w-20 h-20 rounded-full mx-auto mb-3 border-3 border-surface-0 bg-primary-reverse flex items-center justify-center">
                        <i class="pi pi-user text-4xl text-primary"></i>
                    </div>
                    <h3 class="text-xl font-semibold text-primary-contrast mb-1">{{ user.name }}</h3>
                    <p class="text-sm text-primary-contrast opacity-90">{{ user.email }}</p>
                </div>

                <!-- Profile Details -->
                <div class="p-4 space-y-3">
                    <!-- Department -->
                    <div v-if="user.department" class="flex items-center gap-3">
                        <i class="pi pi-briefcase text-muted-color"></i>
                        <div class="flex-1">
                            <p class="text-xs text-muted-color">Department</p>
                            <p class="font-medium">{{ user.department }}</p>
                        </div>
                    </div>

                    <!-- Roles -->
                    <div class="flex items-start gap-3">
                        <i class="pi pi-shield text-muted-color mt-1"></i>
                        <div class="flex-1">
                            <p class="text-xs text-muted-color">Roles</p>
                            <div class="flex flex-wrap gap-1 mt-1">
                                <span 
                                    v-for="role in user.roles" 
                                    :key="role"
                                    class="inline-flex items-center px-2 py-1 rounded text-xs font-medium bg-primary-50 text-primary-700"
                                >
                                    {{ role }}
                                </span>
                            </div>
                        </div>
                    </div>

                    <!-- Permissions -->
                    <div class="flex items-start gap-3">
                        <i class="pi pi-key text-muted-color mt-1"></i>
                        <div class="flex-1">
                            <p class="text-xs text-muted-color">Permissions</p>
                            <div class="flex flex-wrap gap-1 mt-1">
                                <span 
                                    v-for="permission in user.permissions" 
                                    :key="permission"
                                    class="inline-flex items-center px-2 py-1 rounded text-xs font-medium bg-surface-100 text-surface-700"
                                >
                                    {{ permission }}
                                </span>
                            </div>
                        </div>
                    </div>
                </div>

                <!-- Actions -->
                <div class="border-t border-surface-border p-2">
                    <button 
                        type="button" 
                        class="w-full p-3 text-left hover:bg-surface-hover rounded transition-colors flex items-center gap-3"
                    >
                        <i class="pi pi-cog"></i>
                        <span>Settings</span>
                    </button>
                    <button 
                        type="button" 
                        class="w-full p-3 text-left hover:bg-surface-hover rounded transition-colors flex items-center gap-3"
                    >
                        <i class="pi pi-sign-out"></i>
                        <span>Logout</span>
                    </button>
                </div>
            </div>

            <!-- No User -->
            <div v-else class="p-4 text-center">
                <i class="pi pi-user-minus text-4xl text-muted-color mb-3"></i>
                <p class="text-muted-color">Not authenticated</p>
            </div>
        </div>
    </div>
</template>

<style scoped lang="scss">
.layout-profile-menu {
    position: absolute;
    top: calc(100% + 0.5rem);
    right: 0;
    width: 320px;
    background: var(--surface-card);
    border: 1px solid var(--surface-border);
    border-radius: var(--border-radius);
    z-index: 1000;

    .layout-profile-menu-content {
        border-radius: var(--border-radius);
        overflow: hidden;
    }
}

.space-y-3 > * + * {
    margin-top: 0.75rem;
}
</style>
