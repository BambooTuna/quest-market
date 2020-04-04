<template>
  <div class="product-detail">
    <WaitLoading :loading_flag="loadingFlag">
      <h1>{{item.title}}</h1>
      <p class="price">¥ {{item.price}}</p>
      <div class="detail">
        <h3>商品情報</h3>
        <textarea type="text" v-model="item.detail" placeholder="商品詳細" rows="10" readonly></textarea>
      </div>
      <button v-show="item.state === 'open'" @click="purchaseOnClick(item.item_id)">購入申請</button>
      <button v-show="item.state === 'unpaid'" @click="paymentOnClick(item.item_id)">支払い</button>
      <button v-show="item.state === 'sent'" @click="receiptOnClick(item.item_id)">受け取り確認</button>
    </WaitLoading>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Emit, Vue } from 'vue-property-decorator'
import { ContractDetailsResponse } from '@/lib/RestAPIProtocol'
import WaitLoading from '@/components/parts/WaitLoading.vue'

@Component({
  components: {
    WaitLoading
  }
})
export default class ProductDetail extends Vue {
    @Prop()
    private item!: ContractDetailsResponse;

    @Prop()
    private loadingFlag!: boolean

    @Emit()
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    public clickPurchaseEvent (itemId: string): void {
      // do nothing.
    }
    @Emit()
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    public clickPaymentEvent (itemId: string): void {
      // do nothing.
    }
    @Emit()
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    public clickReceiptEvent (itemId: string): void {
      // do nothing.
    }

    purchaseOnClick (itemId: string) {
      this.clickPurchaseEvent(itemId)
    }

    paymentOnClick (itemId: string) {
      this.clickPaymentEvent(itemId)
    }

    receiptOnClick (itemId: string) {
      this.clickReceiptEvent(itemId)
    }
}
</script>

<style scoped>
  .product-detail {
    padding: 0.5em 1em;
    margin: 2em 0;
    font-weight: bold;
    color: #6091d3;/*文字色*/
    background: #FFF;
    border: solid 3px #6091d3;/*線*/
    border-radius: 10px;/*角の丸み*/
  }
  .detail {
    padding: 0.5em 1em;
    margin: 2em 0;
    font-weight: bold;
    color: #000000;/*文字色*/
    background: #FFF;
    border: solid 3px #000000;/*線*/
    border-radius: 2px;/*角の丸み*/
  }
  textarea {
    resize: vertical;
    width: 90%;
    box-sizing: border-box;
    font-size: 16px;
  }
  /* https://saruwakakunw.com/html-css/reference/h-design */
  h1 {
    color: #000000;/*文字色*/
    border-bottom: solid 3px black;
  }
  h3 {
    background: #EEE;/*背景色*/
    padding: 0.5em;/*文字まわり（上下左右）の余白*/
  }
  .price {
    color: #962f10;/*文字色*/
    text-align: left;
    border-radius: 0.5em;/*角丸*/
  }

  button {
    display: inline-block;
    max-width: 180px;
    text-align: left;
    background-color: #ff3308;
    font-size: 16px;
    color: #FFF;
    text-decoration: none;
    font-weight: bold;
    padding: 10px 24px;
    border-radius: 4px;
    border-bottom: 4px solid #d30e01;
  }
  button:active {
    transform: translateY(4px);
    border-bottom: none;
  }
</style>
