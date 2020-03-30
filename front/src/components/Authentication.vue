<template>
  <div class="authentication">
    <WaitLoading :loading_flag="loadingFlag">
      <section class="signup" v-if="!isLogin">
        <h5>新規登録</h5>
        <p><input type="email" v-model="signupMail" placeholder="メールアドレス"></p>
        <p><input type="password" v-model="signupPass" placeholder="パスワード"></p>
        <div class="links">
          <button type="submit" @click="signupEvent()" class="button--signup">新規登録</button>
        </div>
      </section>
      <section class="signin" v-if="!isLogin">
        <h5>ログイン</h5>
        <p><input type="email" v-model="signinMail" placeholder="メールアドレス"></p>
        <p><input type="password" v-model="signinPass" placeholder="パスワード"></p>
        <div class="links">
          <button type="submit" @click="signinEvent()" class="button--signin">ログイン</button>
        </div>
      </section>
      <section class="cooperation" v-if="!isLogin">
        <h5>SNS連携</h5>
        <div class="links">
          <button type="submit" @click="lineCooperationEvent()" class="button--cooperation">Line</button>
        </div>
      </section>
      <section class="logout" v-if="isLogin">
        <h5>ログイン中です</h5>
        <div class="links">
          <button type="submit" @click="logoutEvent()" class="button--logout">ログアウト</button>
        </div>
      </section>
      <section class="slot" v-if="isLogin">
        <slot></slot>
      </section>
    </WaitLoading>
  </div>
</template>

<script lang="ts">
import { Component, Prop, Vue } from 'vue-property-decorator'
import WaitLoading from '@/components/parts/WaitLoading.vue'
import API from '@/lib/RestAPI'

@Component({
  components: {
    WaitLoading
  }
})
export default class Authentication extends Vue {
    private api: API = new API()

    private signupMail?: string = ''
    private signupPass?: string = ''

    private signinMail?: string = ''
    private signinPass?: string = ''

    private isLogin?: boolean = false
    private loadingFlag?: boolean = true

    async created () {
      await this.api
        .health()
        .then(() => {
          this.isLogin = true
        })
        .catch(() => {
          this.isLogin = false
        }).finally(() => {
          this.loadingFlag = false
        })
    }

    init () {
      this.signupMail = ''
      this.signupPass = ''
      this.signinMail = ''
      this.signinPass = ''
    }

    async signupEvent () {
      await this.api
        .signup({ mail: this.signupMail || '', pass: this.signupPass || '' })
        .then(() => {
          this.isLogin = true
        })
        .catch((e: Error) => alert(e.message))
    }

    async signinEvent () {
      await this.api
        .signin({ mail: this.signinMail || '', pass: this.signinPass || '' })
        .then(() => {
          this.isLogin = true
        })
        .catch((e: Error) => alert(e.message))
    }

    async logoutEvent () {
      await this.api
        .logout()
        .then(() => {
          this.isLogin = false
        })
        .catch((e: Error) => alert(e.message))
      this.init()
    }

    async lineCooperationEvent () {
      await this.api.generateLineCooperationUrl()
        .then((redirectUri: string) => {
          window.location.replace(redirectUri)
        })
        .catch((e: Error) => alert(e.message))
    }
}
</script>

<style scoped>
  .signup, .signin, .cooperation, .logout {
    padding: 0.5em 1em;
    margin: 2em 0;
    font-weight: bold;
    color: #6091d3;/*文字色*/
    background: #FFF;
    border: solid 3px #6091d3;/*線*/
    border-radius: 10px;/*角の丸み*/
  }
  input {
    box-sizing: border-box;
    font-size: 16px;
  }

  button {
    display: inline-block;
    max-width: 180px;
    text-align: left;
    background-color: #a6a6a6;
    font-size: 16px;
    color: #FFF;
    text-decoration: none;
    font-weight: bold;
    padding: 10px 24px;
    border-radius: 4px;
    border-bottom: 4px solid #696969;
  }
  button:active {
    transform: translateY(4px);
    border-bottom: none;
  }

  button.button--cooperation {
    display: inline-block;
    max-width: 180px;
    text-align: left;
    background-color: #02d348;
    font-size: 16px;
    color: #FFF;
    text-decoration: none;
    font-weight: bold;
    padding: 10px 24px;
    border-radius: 4px;
    border-bottom: 4px solid #01862e;
  }
  button:active.button--cooperation {
    transform: translateY(4px);
    border-bottom: none;
  }
</style>
