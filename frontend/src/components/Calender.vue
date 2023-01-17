<template>
  <!-- lazy loading tailwind classes -->
  <div class="grid h-full grid-cols-7">
    <div
      v-for="index in maxDaysInMonth"
      @click="changeDate(index)"
      class="flex flex-col justify-center rounded-lg border-2 border-stone-800"
      :class="
        isThisCurrentDate(index)
          ? 'bg-green-200 hover:bg-green-300'
          : 'bg-stone-50 hover:bg-stone-300'
      "
    >
      <p class="text-center text-2xl font-semibold">{{ index }}</p>
    </div>
  </div>
</template>

<script lang="ts">
export default {
  props: {
    currentDateStamp: { type: Number, required: true },
    dateStampToDate: { type: Function, required: true },
  },

  computed: {
    maxDaysInMonth(): number {
      return new Date(
        this.today.getFullYear(),
        this.today.getMonth(),
        0
      ).getDate();
    },
  },

  emits: ["changeDate"],
  methods: {
    changeDate(index: number) {
      this.$emit(
        "changeDate",
        new Date(this.today.getFullYear(), this.today.getMonth(), index)
      );
    },
    isDateEqual(d: Date): boolean {
      let temp: Date = this.dateStampToDate(this.currentDateStamp);
      return (
        d.getDate() == temp.getDate() &&
        d.getMonth() == temp.getMonth() &&
        d.getFullYear() == temp.getFullYear()
      );
    },
    isThisCurrentDate(index: number): boolean {
      console.log;
      return this.isDateEqual(
        new Date(this.today.getFullYear(), this.today.getMonth(), index)
      );
    },
  },

  data() {
    return {
      today: new Date(),
    };
  },
};
</script>
