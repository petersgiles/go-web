<script setup>
import { ProductService } from '@/service/ProductService';
import { onMounted, ref } from 'vue';

const orders = ref([]);

function formatCurrency(value) {
    return value.toLocaleString('en-US', { style: 'currency', currency: 'USD' });
}

function getStatusSeverity(status) {
    switch (status) {
        case 'DELIVERED':
            return 'success';
        case 'PENDING':
            return 'warn';
        case 'CANCELLED':
            return 'danger';
        case 'RETURNED':
            return 'info';
        default:
            return null;
    }
}

onMounted(() => {
    ProductService.getProductsWithOrders().then((data) => {
        // Flatten all orders from all products
        const allOrders = [];
        data.forEach(product => {
            if (product.orders) {
                product.orders.forEach(order => {
                    allOrders.push({
                        ...order,
                        productName: product.name,
                        productImage: product.image
                    });
                });
            }
        });
        // Sort by date descending and take most recent
        allOrders.sort((a, b) => new Date(b.date) - new Date(a.date));
        orders.value = allOrders;
    });
});
</script>

<template>
    <div class="card">
        <div class="font-semibold text-xl mb-4">Recent Sales</div>
        <DataTable :value="orders" :rows="5" :paginator="true" responsiveLayout="scroll">
            <Column field="customer" header="Customer" :sortable="true" style="width: 25%"></Column>
            <Column field="productName" header="Product" :sortable="true" style="width: 25%"></Column>
            <Column field="date" header="Date" :sortable="true" style="width: 20%"></Column>
            <Column field="amount" header="Amount" :sortable="true" style="width: 15%">
                <template #body="slotProps">
                    {{ formatCurrency(slotProps.data.amount) }}
                </template>
            </Column>
            <Column field="status" header="Status" :sortable="true" style="width: 15%">
                <template #body="slotProps">
                    <Tag :value="slotProps.data.status" :severity="getStatusSeverity(slotProps.data.status)" />
                </template>
            </Column>
        </DataTable>
    </div>
</template>
