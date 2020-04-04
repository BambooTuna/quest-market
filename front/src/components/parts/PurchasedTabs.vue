<template>
  <div class="purchased-tabs">
    <WaitLoading :loadingFlag="loadingFlag">
    </WaitLoading>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import WaitLoading from '@/components/parts/WaitLoading.vue'

@Component({
  components: {
    WaitLoading
  }
})
export default class PurchasedTabs extends Vue {
    // @Prop()
    // private items!: Array<ProductTransaction>;

    @Prop()
    private loadingFlag!: boolean
}
</script>

<style scoped>
  .purchased-tabs {
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
