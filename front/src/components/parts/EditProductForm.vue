<template>
  <div class="edit_product_form">
    <WaitLoading :loading_flag="loadingFlag">
      <table>
        <tbody>
        <tr>
          <th colspan="2" class="table-title">商品の説明</th>
        </tr>
        <tr>
          <th>商品名</th>
          <td><input type="text" v-model="item.productTitle" placeholder="タイトル" class="title-input"></td>
        </tr>
        <tr>
          <th>詳細説明</th>
          <td><textarea type="text" v-model="item.productDetail" placeholder="商品詳細" rows="5"></textarea></td>
        </tr>
        <tr>
          <th>価格</th>
          <td><input type="number" v-model="item.requestPrice" placeholder="価格" class="price-input"></td>
        </tr>
        </tbody>
      </table><br><br><br>
      <button @click="onClick('open')">出品</button><button @click="onClick('draft')">下書き</button><br>
      <button v-show="!isNew" @click="onClick('closed')">削除</button>
    </WaitLoading>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Emit, Vue } from 'vue-property-decorator'
import { ProductDetailRequest, ProductDetailResponse, StateEnum } from '@/lib/RestAPIProtocol'
import WaitLoading from '@/components/parts/WaitLoading.vue'

@Component({
  components: {
    WaitLoading
  }
})
export default class EditProductForm extends Vue {
    @Prop()
    private isNew!: boolean

    @Prop()
    private loadingFlag!: boolean

    @Prop()
    private item!: ProductDetailResponse;

    @Emit()
    // eslint-disable-next-line @typescript-eslint/no-unused-vars
    public clickEvent (data: ProductDetailRequest): void {
      // do nothing.
    }

    onClick (state: StateEnum) {
      this.clickEvent({
        title: this.item.productTitle,
        detail: this.item.productDetail,
        price: this.item.requestPrice,
        state: state
      })
    }
}
</script>

<style scoped>
  .edit_product_form {
    padding: 0.5em 1em;
    margin: 2em 0;
    font-weight: bold;
    color: #6091d3;/*文字色*/
    background: #FFF;
    border: solid 3px #6091d3;/*線*/
    border-radius: 10px;/*角の丸み*/
  }
  .title-input {
    width: 90%;
  }
  input {
    box-sizing: border-box;
    font-size: 16px;
  }
  textarea {
    resize: vertical;
    width: 90%;
    box-sizing: border-box;
    font-size: 16px;
  }
  .price-input {
    width: 90%;
  }

  table {
    border-collapse:  collapse; /* セルの線を重ねる */
    width: 100%;
  }
  th,td {
    border: solid 1px #d1d1d1;          /* 枠線指定 */
  }
  th {
    background-color: #f2f2f2; /* 背景色指定 */
    color: #000000;               /* 文字色指定 */
    font-weight:  normal;       /* 文字の太さ指定 */
  }
  td {
    text-align: left;
    background-color: #fbfbfb; /* 背景色指定 */
  }
  th.table-title {
    background-color: #636363; /* 背景色指定 */
    color: #f2f2f2;               /* 文字色指定 */
  }
  /*td.table-button {*/
  /*  text-align: center;*/
  /*  background-color: #fafbff; !* 背景色指定 *!*/
  /*}*/
  /*td.table-delete-button {*/
  /*  text-align: center;*/
  /*  background-color: #fafbff; !* 背景色指定 *!*/
  /*}*/

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
