<template>
  <div class="public-products-table">
        <ProductsTable :items="productList" :privateMode="false" :loadingFlag="loadingFlag"></ProductsTable>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import ProductsTable from '@/components/parts/ProductsTable.vue'
import API from '@/lib/RestAPI'
import { DisplayLimit, ProductDetailResponse } from '@/lib/RestAPIProtocol'

@Component({
  components: {
    ProductsTable
  }
})
export default class PublicProductsTable extends Vue {
    private api: API = new API()
    private productList: Array<ProductDetailResponse> = []
    private loadingFlag?: boolean = true

    @Prop()
    private params!: DisplayLimit

    async created () {
      await this.api
        .getProducts(this.params)
        .then(r => {
          this.productList = r
        })
        .catch((e: Error) => alert(e.message))
        .finally(() => {
          this.loadingFlag = false
        })
    }
}
</script>
