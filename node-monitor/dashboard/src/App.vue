<template>
  <a-layout class="layout">
    <div>
      <a-layout-header>
        <a-menu v-model:selected-keys="selectedKeys" @menuItemClick="onClickMenuItem" mode='horizontal'>
          <!--<a-menu-item key="0" :style="{ padding: 0, marginRight: '38px' }" disabled>
            <div :style="{
              width: '30px',
              height: '30px',
            }">
              <a-avatar :size="30" :style="{ backgroundColor: '#3370ff' }">
                node-monitor
              </a-avatar>
            </div>
          </a-menu-item>-->
          <a-menu-item key="0_2">
            <IconStorage />
            Container
          </a-menu-item>
        </a-menu>
      </a-layout-header>
      <a-layout-content style="padding: 12px 24px;">
        <RouterView />
      </a-layout-content>
      <a-layout-footer style="margin: 24px 36px;">
      </a-layout-footer>
    </div>
  </a-layout>
</template>
<script>
import { defineComponent, ref } from 'vue'
import '@arco-design/web-vue/es/message/style/css.js'

import {
  IconHome,
  IconStorage,
  IconSettings,
  IconImport,
  IconGithub,
} from '@arco-design/web-vue/es/icon'

import { useRouter } from 'vue-router'

export default defineComponent({
  components: {
    IconHome,
    IconStorage,
    IconSettings,
    IconImport,
    IconGithub,
  },
  data: () => {
    return {
      formLogin: {},
      selectedKeys: ['0_1'],
    }
  },
  setup() {
    const isLogin = ref(true)
    const isAdmin = ref(true)

    const router = useRouter()
    return {
      onClickMenuItem: (key) => {
        switch (key) {
          case '0_2':
            router.push({ path: '/container' })
            break
          default:
            break
        }
      },
      isLogin,
      isAdmin,
    }
  },
  watch: {
    '$route'(to) {
      if (to.path === '/container') this.selectedKeys[0] = '0_2'
    }
  },
  mounted() {
  },
})
</script>
<style scoped>
.layout {
  height: 100%;
}

.loginContainer {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.loginDiv {
  width: 40%;
  margin-bottom: 6%;
}
</style>
