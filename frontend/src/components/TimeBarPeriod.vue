<template>
  <div
    class="flex w-full place-content-center items-center border-2"
    :style="{
      height: h + 'px',
      borderColor: getTypeColor(type, 2),
      backgroundColor: hover ? getTypeColor(type, 1) : getTypeColor(type, 0),
      color: getTypeColor(type, 2),
    }"
    @mouseover="hover = true"
    @mouseleave="hover = false"
  >
    <p v-if="isShown || hover" class="text-center font-mono font-semibold">
      {{ text }}
    </p>
  </div>
</template>

<script lang="ts">
export default {
  props: {
    h: { type: Number, required: true },
    text: { type: String, required: true },
    type: { type: String, required: true },
    getTypeColor: { type: Function, required: true },
  },
  data() {
    return {
      hover: false,
    };
  },
  computed: {
    isShown(): boolean {
      return this.h > 14;
    },
  },
};
</script>
