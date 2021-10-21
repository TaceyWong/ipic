<template>
  <section class="bg-gray-800 overflow-y-auto p-10 w-full">
    <header class="flex items-center justify-between w-full">
      <h1 class="font-medium text-2xl text-purple-400">关于</h1>
      <BtnClose @click.native="closeView" color="purple" />
    </header>
    <div class="flex flex-wrap items-center justify-center pt-16 w-full">
      <div class="flex items-center justify-center logo w-full h-64"></div>
      <!-- info -->
      <div class="text-center text-gray-100 w-full">
        <p class="mb-1">{{ pkg.name }} {{ pkg.version }}</p>
        <p class="mb-1">
          © 2021
          <span
            class="cursor-pointer hover:text-purple-400 ta-color-slow"
            @click="openLink('https://github.com/TaceyWong')"
            >{{ pkg.author.name }}</span
          >
        </p>
        <p class="mb-1">
          <span
            class="cursor-pointer hover:text-purple-400 ta-color-slow"
            @click="
              openLink('https://github.com/TaceyWong/ipic/blob/main/LICENSE')
            "
            >协议</span
          >
          和
          <span
            class="cursor-pointer hover:text-purple-400 ta-color-slow"
            @click="openLink('https://github.com/TaceyWong/ipic')"
            >文档</span
          >
        </p>
      </div>
    </div>
  </section>
</template>

<script>
import BtnClose from './BtnClose'
import pkg from './../../package.json'

export default {
  name: 'About',

  components: { BtnClose },

  data() {
    return {
      pkg
    }
  },

  methods: {
    closeView() {
      this.$emit('close-view')
    },

    openLink(u) {
      window.wails.Browser.OpenURL(u)
    }
  }
}
</script>

<style scoped>
.logo {
  background-image: url('./../assets/images/logo.jpg');
  background-position: center;
  background-repeat: no-repeat;
  background-size: contain;
}

span {
  box-shadow: 0 1px #d690ff;
}
</style>
