<template>
  <div class="exhibition">
    <Authentication>
      <EditProductForm :item="item" @click-event="clickEvent" :isNew="true"></EditProductForm>
    </Authentication>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import Authentication from '@/components/Authentication.vue'
import EditProductForm from '@/components/parts/EditProductForm.vue'
import PrivateProductsTable from '@/components/PrivateProductsTable.vue'
import API from '@/lib/RestAPI'
import { ProductDetailRequest, ProductDetailResponse } from '@/lib/RestAPIProtocol'

@Component({
  components: {
    Authentication, EditProductForm, PrivateProductsTable
  }
})
export default class Exhibition extends Vue {
  private api: API = new API()

  private item: ProductDetailResponse = {
    id: '',
    productTitle: '',
    productDetail: '',
    requestPrice: 0,
    presenterId: '',
    state: 'open'
  }

  clickEvent (data: ProductDetailRequest) {
    this.api.postProduct(data)
      .then(() => {
        alert('出品完了')
      })
      .catch((e: Error) => alert(e.message))
  }
}
</script>

<style scoped>

</style>
