<template>
  <div class="product-detail-view">
    <ProductDetail v-show="!isEditMode()" :item="item" :loadingFlag="loadingFlag"></ProductDetail>
    <EditProductForm v-show="isEditMode()" :item="item" @click-event="clickEvent" :isNew="false" :loadingFlag="loadingFlag"></EditProductForm>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import ProductDetail from '@/components/parts/ProductDetail.vue'
import EditProductForm from '@/components/parts/EditProductForm.vue'
import API from '@/lib/RestAPI'
import { ProductDetailRequest, ProductDetailResponse } from '@/lib/RestAPIProtocol'

@Component({
  components: {
    ProductDetail, EditProductForm
  }
})
export default class ProductDetailView extends Vue {
    private api: API = new API()
    private item: ProductDetailResponse = {
      id: '',
      productTitle: '',
      productDetail: '',
      requestPrice: 0,
      presenterId: '',
      state: 'draft'
    }

    public loadingFlag?: boolean = true
    private productId!: string

    isEditMode (): boolean {
      return this.$route.query.mode === 'edit'
    }

    async created () {
      this.productId = this.$route.params.id
      if (this.isEditMode()) {
        await this.api.getMyProductDetail(this.productId)
          .then((res: ProductDetailResponse) => {
            this.item = res
          })
          .catch(() => alert('見つかりません'))
          .finally(() => {
            this.loadingFlag = false
          })
      } else {
        await this.api.getProductDetail(this.productId)
          .then((res: ProductDetailResponse) => {
            this.item = res
          })
          .catch(() => alert('見つかりません'))
          .finally(() => {
            this.loadingFlag = false
          })
      }
    }

    clickEvent (data: ProductDetailRequest) {
      this.api.editProduct(this.productId, data)
        .then(() => alert('編集完了'))
        .catch((e: Error) => alert(e.message))
    }
}
</script>

<style scoped>

</style>
