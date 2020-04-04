<template>
  <div class="private-products-table">
        <ProductsTable :items="items" :privateMode="true" :loadingFlag="loadingFlag"></ProductsTable>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import ProductsTable from '@/components/parts/ProductsTable.vue'
import API from '@/lib/RestAPI'
import { StateDisplayLimit, ContractDetailsResponse } from '@/lib/RestAPIProtocol'

@Component({
  components: {
    ProductsTable
  }
})
export default class PrivateProductsTable extends Vue {
    private api: API = new API()
    private items: Array<ContractDetailsResponse> = []
    private loadingFlag?: boolean = true

    @Prop()
    private params!: StateDisplayLimit

    async created () {
      await this.api
        .getMyProducts(this.params)
        .then(r => {
          this.items = r
        })
        .finally(() => {
          this.loadingFlag = false
        })
    }
}
</script>
