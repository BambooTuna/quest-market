<template>
  <div class="product-detail-view">
    <ProductDetail v-show="!isEditMode()" @click-purchase-event="clickPurchaseEvent" @click-payment-event="clickPaymentEvent" @click-receipt-event="clickReceiptEvent" :item="item" :loadingFlag="loadingFlag"></ProductDetail>
    <EditProductForm v-show="isEditMode()" :item="item" @click-event="clickEvent" :isNew="false" :loadingFlag="loadingFlag"></EditProductForm>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import ProductDetail from '@/components/parts/ProductDetail.vue'
import EditProductForm from '@/components/parts/EditProductForm.vue'
import API from '@/lib/RestAPI'
import { ProductDetailRequest, ContractDetailsResponse } from '@/lib/RestAPIProtocol'

@Component({
  components: {
    ProductDetail, EditProductForm
  }
})
export default class ProductDetailView extends Vue {
    private api: API = new API()
    private item: ContractDetailsResponse = {
      // eslint-disable-next-line @typescript-eslint/camelcase
      item_id: '',
      title: '',
      detail: '',
      price: 0,
      // eslint-disable-next-line @typescript-eslint/camelcase
      seller_account_id: '',
      // eslint-disable-next-line @typescript-eslint/camelcase
      updated_at: '',
      state: 'draft',
      accessor: 'general'
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
          .then((res: ContractDetailsResponse) => {
            this.item = res
          })
          .catch(() => alert('見つかりません'))
          .finally(() => {
            this.loadingFlag = false
          })
      } else {
        await this.api.getProductDetail(this.productId)
          .then((res: ContractDetailsResponse) => {
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

    clickPurchaseEvent (itemId: string) {
      this.api.purchaseItem(itemId)
        .then(() => {
          alert('購入申請完了')
        })
        .catch((e: Error) => alert(e.message))
    }

    clickPaymentEvent (itemId: string) {
      this.api.paymentForItem(itemId)
        .then(() => {
          alert('支払い完了')
        })
        .catch((e: Error) => alert(e.message))
    }

    clickReceiptEvent (itemId: string) {
      this.api.receiptOfItem(itemId)
        .then(() => {
          alert('受け取り確認完了')
        })
        .catch((e: Error) => alert(e.message))
    }
}
</script>

<style scoped>

</style>
