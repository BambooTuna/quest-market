<template>
  <div class="private-products-table">
        <ProductsTable :items="productList" :privateMode="true" :loadingFlag="loadingFlag"></ProductsTable>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import ProductsTable from '@/components/parts/ProductsTable.vue'
import API from '@/lib/RestAPI'
import { StateDisplayLimit, ProductDetailResponse } from '@/lib/RestAPIProtocol'

@Component({
  components: {
    ProductsTable
  }
})
export default class PrivateProductsTable extends Vue {
    private api: API = new API()
    private productList: Array<ProductDetailResponse> = []
    private loadingFlag?: boolean = true

    @Prop()
    private params!: StateDisplayLimit

    async created () {
      await this.api
        .getMyProducts(this.params)
        .then(r => {
          this.productList = r
        })
        .finally(() => {
          this.loadingFlag = false
        })
    }
}
</script>
