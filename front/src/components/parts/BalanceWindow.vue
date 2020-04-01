<template>
  <div class="balance-window">
    <WaitLoading :loading_flag="loadingFlag">
      <p class="price">残高: ¥ {{item.balance}}</p>
    </WaitLoading>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import { Balance } from '@/lib/RestAPIProtocol'
import WaitLoading from '@/components/parts/WaitLoading.vue'

@Component({
  components: {
    WaitLoading
  }
})
export default class BalanceWindow extends Vue {
    @Prop()
    private item!: Balance;

    @Prop()
    private loadingFlag!: boolean
}
</script>

<style scoped>
  .balance-window {
    padding: 0.5em 1em;
    margin: 2em 0;
    font-weight: bold;
    color: #6091d3;/*文字色*/
    background: #FFF;
    border: solid 3px #6091d3;/*線*/
    border-radius: 10px;/*角の丸み*/
  }

  .price {
    color: #962f10;/*文字色*/
    text-align: left;
    border-radius: 0.5em;/*角丸*/
  }
</style>
