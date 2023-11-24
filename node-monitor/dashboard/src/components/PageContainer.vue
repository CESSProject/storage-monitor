<template>
  <div>
    <a-modal v-model:visible="visibleNew" title="New Conatiner" @before-ok="handleBeforeOkNew">
    </a-modal>

    <a-page-header :show-back="false">
      <template #extra>
        <a-space>
          <a-button @click="handleClickRefresh" type="primary">
            <template #icon>
              <icon-sync />
            </template>
            <template #default>Refresh</template>
          </a-button>
          <a-button @click="handleClickHelp" type="primary" status="warning">
            <template #icon>
              <icon-exclamation-circle />
            </template>
            <template #default>Need Help?</template>
          </a-button>
        </a-space>
      </template>
    </a-page-header>

    <a-card v-for="info in containerList" :key="info.name" :style="{ width: 'auto', margin: '8px 12px' }" :title="info.name"
      hoverable>
      <template #extra>

      </template>
      <a-descriptions :data="[
        ...(() => {
          return Object.entries(info.data).map((data) => {
            return {
              label: data[0],
              value: data[1],
            }
          })
        })(),
      ]" :size="'medium'" :column="2" />
    </a-card>
  </div>
</template>

<script>
import { defineComponent } from 'vue'
import { Modal, Message } from '@arco-design/web-vue'
import '@arco-design/web-vue/es/message/style/css.js'
import '@arco-design/web-vue/es/modal/style/css.js'

import axios from 'axios'
import {
  IconPlusCircle,
  IconExclamationCircle,
  IconRefresh,
  IconDelete,
  IconSync,
} from '@arco-design/web-vue/es/icon'

export default defineComponent({
  components: {
    IconPlusCircle,
    IconExclamationCircle,
    IconRefresh,
    IconDelete,
    IconSync,
  },
  data: () => {
    return {
      visibleNew: false,
      containerList: [],
    }
  },
  computed: {
  },
  methods: {
    getMeta() {
      return axios.get('/miners').then((res) => {
        if (res.status === 200) {
          const data = res.data
          this.containerList = data.map((item) => {
            return {
              name: item.container_info.name,
              containerId: item.container_info.id,
              stakingAccount:item.miner_metadata.staking_account,
              data: {
                'ContainerState': item.container_info.state,
                'Image': item.container_info.image,
                'ServerName': item.miner_metadata.name,
                "ContainerId":item.container_info.id,
                'ServerState': item.miner_metadata.state,
                'PeerId': item.miner_metadata.peer_id,
                'ValidatedSpace': (item.miner_metadata.validated_space/1024/1024/1024).toFixed(2)+" GiB",
                'UsedSpace': (item.miner_metadata.used_space/1024/1024/1024).toFixed(2)+" GiB",
                'LockedSpace': (item.miner_metadata.locked_space/1024/1024/1024).toFixed(2)+" GiB",
                'StakingAmount': item.miner_metadata.staking_amount,
                'CPUUsage': item.container_status.cpu + "%",
                'MemoryPercent': item.container_status.memory+ "%",
                'MemoryUsage': (item.container_status.mem_usage/1024/1024/1024).toFixed(2)+" GiB",
                'EarningsAccount': item.miner_metadata.earnings_account,
              },
            }
          })
        }
      }).catch((err) => {
        console.log(err)
      })
    },
    handleClickNew() {
      this.visibleNew = true
      this.formNew.server = this.serverList[0]
      this.formNew.image = this.imageList[0].name
    },
    handleClickRefresh() {
      this.getMeta()
      Message.success('Refreshed')
    },
    handleClickHelp() {
      Modal.info({
        title: 'Contact Admin',
        content: 'Should you require any assistance, kindly reach out to the administrator at jiuqu3239@gmail.com'
      })
    },
    modalRestartContainer(containerId) {
      Modal.confirm({
        title: 'Restart Container',
        content: 'Are you sure to restart this container?',
        escToClose: false,
        maskClosable: false,
        onBeforeOk: () => {
          return this.restartContainer(containerId).then(() => this.modalWait())
        },
      })
    },
    modalRemoveContainer(containerId) {
      Modal.confirm({
        title: 'Remove Container',
        content: 'Are you sure to remove this container?',
        escToClose: false,
        maskClosable: false,
        onBeforeOk: () => {
          return this.removeContainer(containerId).then(() => this.modalWait())
        },
      })
    },
    modalWait() {
      const modal = Modal.confirm({
        title: 'Syncing',
        content: 'It might take 20 seconds, please wait...',
        escToClose: false,
        maskClosable: false,
        hideCancel: true,
        okLoading: true,
      })

      axios.get('/container/wait').then((res) => {
        if (res.data.status === 200) {
          this.getMeta()
          Message.success('Data synced')
          modal.close()
        }
      }).catch((err) => {
        console.log(err)
        this.getMeta()
        Message.error('Data synced with error')
        modal.close()
      })

      return true
    },
    handleBeforeOkNew() {
      return axios.get('/container/create', {
        params: {
          server: this.formNew.server,
          image: this.formNew.image,
        },
      }).then((res) => {
        if (res.data.status === 200) {
          // console.log(res.data, res.data.data)
          Message.success('Created')
        }
      }).catch((err) => {
        console.log(err)
        if (err.response.data.status === 400) {
          Message.error(err.response.data.data.msg)
        }
      }).then(() => this.modalWait())
    },
    restartContainer(containerId) {
      return axios.get('/container/restart', {
        params: {
          containerId,
        },
      }).then((res) => {
        if (res.data.status === 200) {
          Message.success('Restarted')
        }
      }).catch((err) => {
        console.log(err)
      })
    },
    removeContainer(containerId) {
      return axios.get('/container/stopAndRemove', {
        params: {
          containerId,
        },
      }).then((res) => {
        if (res.data.status === 200) {
          Message.success('Removed')
        }
      }).catch((err) => {
        console.log(err)
      })
    },
  },
  mounted() {
    this.getMeta()
  },
})
</script>

<style scoped></style>
