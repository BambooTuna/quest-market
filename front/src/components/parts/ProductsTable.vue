<template>
  <div class="products_table">
    <WaitLoading :loading_flag="loadingFlag">
      <ul>
        <li v-for="row in items" :key="row.id">
          <h2><router-link :to=" '/product/' + row.id + (privateMode ? '?mode=edit' : '')">{{row.productTitle}}</router-link></h2>
          <p class="price">¥ {{row.requestPrice}}</p>
          <p class="state" v-show="privateMode">{{ (row.state === 'draft' ? '下書き' : '出品中') }}</p>
        </li>
      </ul>
    </WaitLoading>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Emit, Vue } from 'vue-property-decorator'
import { ProductDetailResponse } from '@/lib/RestAPIProtocol'
import WaitLoading from '@/components/parts/WaitLoading.vue'

@Component({
  components: {
    WaitLoading
  }
})
export default class ProductsTable extends Vue {
    @Prop()
    private items!: Array<ProductDetailResponse>;

    @Prop()
    private privateMode!: boolean

    @Prop()
    private loadingFlag?: boolean = true
}
</script>

<style scoped>
  ul, ol {
    padding: 0;
    position: relative;
  }
  ul li, ol li {
    color: #000000;
    border-left: solid 6px #000000;/*左側の線*/
    border-right: solid 6px #000000;/*左側の線*/
    border-bottom: solid 6px #000000;/*左側の線*/
    border-top: solid 6px #000000;/*左側の線*/
    background: rgba(241, 248, 255, 0.26);/*背景色*/
    margin-bottom: 3px;/*下のバーとの余白*/
    line-height: 1.5;
    padding: 0.5em;
    list-style-type: none!important;/*ポチ消す*/
  }
  h2 {
    text-align: left;
  }
  .price {
    color: #962f10;/*文字色*/
    text-align: right;
    border-radius: 0.5em;/*角丸*/
  }
  .state {
    color: #000000;/*文字色*/
    text-align: right;
  }
</style>
