<template>
  <div class="purchased-window">
    <input id="tab1" type="radio" name="TAB" class="tab-switch" checked="checked" /><label class="tab-label" for="tab1">売ったもの</label>
    <div class="tab-content">
      <ProductsTable :items="sellItems()" :privateMode="true" :loadingFlag="loadingFlag"></ProductsTable>
    </div>
    <input id="tab2" type="radio" name="TAB" class="tab-switch" /><label class="tab-label" for="tab2">買ったもの</label>
    <div class="tab-content">
      <ProductsTable :items="buyItems()" :loadingFlag="loadingFlag"></ProductsTable>
    </div>
    <input id="tab3" type="radio" name="TAB" class="tab-switch" /><label class="tab-label" for="tab3">下書き</label>
    <div class="tab-content">
      <ProductsTable :items="draftItems()" :privateMode="true" :loadingFlag="loadingFlag"></ProductsTable>
    </div>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import ProductsTable from '@/components/parts/ProductsTable.vue'
import API from '@/lib/RestAPI'
import { ContractDetailsResponse, DisplayLimit } from '@/lib/RestAPIProtocol'

@Component({
  components: {
    ProductsTable
  }
})
export default class PurchasedWindow extends Vue {
    private api: API = new API()
    private items: Array<ContractDetailsResponse> = []
    private loadingFlag?: boolean = true

    @Prop()
    private params!: DisplayLimit

    created (): void {
      this.api.getMyProducts(this.params)
        .then((items: Array<ContractDetailsResponse>) => {
          this.items = items
          this.loadingFlag = false
        })
    }

    sellItems (): Array<ContractDetailsResponse> {
      return this.items.filter((a: ContractDetailsResponse) => (a.accessor === 'seller' && a.state !== 'draft'))
    }

    buyItems (): Array<ContractDetailsResponse> {
      return this.items.filter((a: ContractDetailsResponse) => a.accessor === 'buyer')
    }

    draftItems (): Array<ContractDetailsResponse> {
      return this.items.filter((a: ContractDetailsResponse) => a.accessor === 'seller' && a.state === 'draft')
    }
}
</script>

<style scoped>
  .purchased-window {
    display: flex;
    flex-wrap: wrap;

    padding: 0.5em 1em;
    margin: 2em 0;
    font-weight: bold;
    color: #6091d3;/*文字色*/
    background: #FFF;
    border: solid 3px #6091d3;/*線*/
    border-radius: 10px;/*角の丸み*/
  }
  .tab-label {
    color: White;
    background: LightGray;
    margin-right: 5px;
    padding: 3px 12px;
    order: -1;
  }
  .tab-content {
    width: 100%;
    display: none;
  }
  /* アクティブなタブ */
  .tab-switch:checked+.tab-label {
    background: DeepSkyBlue;
  }
  .tab-switch:checked+.tab-label+.tab-content {
    display: block;
  }
  /* ラジオボタン非表示 */
  .tab-switch {
    display: none;
  }
</style>
