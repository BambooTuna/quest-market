<template>
  <div class="purchased-window">
<!--    <input id="tab1" type="radio" name="TAB" class="tab-switch" checked="checked" /><label class="tab-label" for="tab1">ボタン 1</label>-->
<!--    <div class="tab-content">-->
<!--      <PurchasedTabs :items="items" :loadingFlag="loadingFlag"></PurchasedTabs>-->
<!--    </div>-->
<!--    <input id="tab2" type="radio" name="TAB" class="tab-switch" /><label class="tab-label" for="tab2">ボタン 2</label>-->
<!--    <div class="tab-content">-->
<!--      コンテンツ 2-->
<!--    </div>-->
<!--    <input id="tab3" type="radio" name="TAB" class="tab-switch" /><label class="tab-label" for="tab3">ボタン 3</label>-->
<!--    <div class="tab-content">-->
<!--      コンテンツ 3-->
<!--    </div>-->
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import PurchasedTabs from '@/components/parts/PurchasedTabs.vue'
import API from '@/lib/RestAPI'

@Component({
  components: {
    PurchasedTabs
  }
})
export default class PurchasedWindow extends Vue {
    // private api: API = new API()
    // private items: Array<ProductTransaction> = []
    // private loadingFlag?: boolean = true
    //
    // created (): void {
    //   this.api.getMyProductTransaction()
    //     .then((items: Array<ProductTransaction>) => {
    //       this.items = items
    //       this.loadingFlag = false
    //     })
    // }
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
