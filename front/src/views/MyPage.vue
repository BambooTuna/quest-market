<template>
  <div class="my-page">
    <Authentication>
      <BalanceWindow :item="item" :loadingFlag="loadingFlag"></BalanceWindow>
      <PurchasedWindow :params="this.$route.query"></PurchasedWindow>
    </Authentication>
  </div>
</template>

<script lang="ts">
import { Component, Vue } from 'vue-property-decorator'
import Authentication from '@/components/Authentication.vue'
import PrivateProductsTable from '@/components/PrivateProductsTable.vue'
import BalanceWindow from '@/components/parts/BalanceWindow.vue'
import PurchasedWindow from '@/components/PurchasedWindow.vue'
import API from '@/lib/RestAPI'
import { Balance } from '@/lib/RestAPIProtocol'

@Component({
  components: {
    Authentication, PrivateProductsTable, BalanceWindow, PurchasedWindow
  }
})

export default class MyPage extends Vue {
    private api: API =new API()
    private item!: Balance
    private loadingFlag?: boolean = true

    async created () {
      await this.api.getBalance()
        .then((balance: Balance) => {
          this.item = balance
          this.loadingFlag = false
        })
    }
}
</script>

<style scoped>

</style>
