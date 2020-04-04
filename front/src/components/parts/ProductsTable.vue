<template>
  <div class="products_table">
    <WaitLoading :loading_flag="loadingFlag">
      <ul>
        <li v-for="row in items" :key="row.item_id">
          <h2><router-link :to=" '/product/' + row.item_id + ((privateMode && (row.state === 'open' || row.state === 'draft')) ? '?mode=edit' : '')">{{row.title}}</router-link></h2>
          <p class="price">¥ {{row.price}}</p>
          <p class="state">{{ stateMessage(row.state) }}</p>
        </li>
      </ul>
    </WaitLoading>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Emit, Vue } from 'vue-property-decorator'
import { ContractDetailsResponse, StateEnum } from '@/lib/RestAPIProtocol'
import WaitLoading from '@/components/parts/WaitLoading.vue'

@Component({
  components: {
    WaitLoading
  }
})
export default class ProductsTable extends Vue {
    @Prop()
    private items!: Array<ContractDetailsResponse>;

    @Prop()
    private privateMode!: boolean

    @Prop()
    private loadingFlag?: boolean = true

    stateMessage (state: StateEnum): string {
      let r = '???'
      switch (state) {
        case 'open':
          r = '出品中'
          break
        case 'sold':
          r = '売り切れ'
          break
        case 'draft':
          r = '下書き'
          break
        case 'unpaid':
          r = '支払い待ち'
          break
        case 'sent':
          r = '受け取り待ち'
          break
        case 'complete':
          r = '取引終了'
          break
        case 'deleted':
          r = '削除済み'
          break
      }
      return r
    }
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
