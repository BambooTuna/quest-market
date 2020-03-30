<template>
  <div class="product-detail">
    <WaitLoading :loading_flag="loadingFlag">
      <h1>{{item.productTitle}}</h1>
      <p class="price">¥ {{item.requestPrice}}</p>
      <div class="detail">
        <h3>商品情報</h3>
        <textarea type="text" v-model="item.productDetail" placeholder="商品詳細" rows="10" readonly></textarea>
      </div>
    </WaitLoading>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { ProductDetailResponse } from '@/lib/RestAPIProtocol'
import WaitLoading from '@/components/parts/WaitLoading.vue'

@Component({
  components: {
    WaitLoading
  }
})
export default class ProductDetail extends Vue {
    @Prop()
    private item!: ProductDetailResponse;

    @Prop()
    private loadingFlag!: boolean
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
</style>
