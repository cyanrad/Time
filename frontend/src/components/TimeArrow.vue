<template>
  <div v-if="h > 14" class="absolute" :style="{ top: displace - 15 + 'px' }">
    <img class="mr-1 inline" src="@/assets/arrow.png" width="12" height="12" />
    <p
      :style="{ color: getTypeColor(type, 1) }"
      class="inline font-mono text-sm"
    >
      {{ time }}
    </p>
  </div>
  <div v-else class="absolute" :style="{ top: displace - 15 + 'px' }">
    <img class="mr-1 inline" src="@/assets/arrow.png" width="6" height="6" />
    <p
      :style="{ color: getTypeColor(type, 1) }"
      class="inline font-mono text-stone-50"
      style="font-size: 10px"
    >
      {{ time }}
    </p>
  </div>
</template>

<script lang="ts">
export default {
  props: {
    time: String,
    h: { type: Number, required: true },
    displace: { type: Number, required: true },
    type: { type: String, required: true },
    getTypeColor: { type: Function, required: true },
  },
};
</script>
