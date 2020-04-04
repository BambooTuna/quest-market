<template>
  <div class="public-products-table">
        <ProductsTable :items="items" :privateMode="false" :loadingFlag="loadingFlag"></ProductsTable>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import ProductsTable from '@/components/parts/ProductsTable.vue'
import API from '@/lib/RestAPI'
import { DisplayLimit, ContractDetailsResponse } from '@/lib/RestAPIProtocol'

@Component({
  components: {
    ProductsTable
  }
})
export default class PublicProductsTable extends Vue {
    private api: API = new API()
    private items: Array<ContractDetailsResponse> = []
    private loadingFlag?: boolean = true

    @Prop()
    private params!: DisplayLimit

    async created () {
      await this.api
        .getProducts(this.params)
        .then(r => {
          this.items = r
        })
        .catch((e: Error) => alert(e.message))
        .finally(() => {
          this.loadingFlag = false
        })
    }
}
</script>
