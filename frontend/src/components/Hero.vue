<template>
  <section class="relative min-h-screen flex items-center justify-center overflow-hidden">
    <!-- Animated Background -->
    <div class="absolute inset-0 mesh-gradient opacity-90"></div>
    <div
      class="absolute inset-0 bg-gradient-to-br from-primary/20 via-secondary/20 to-accent/20"
    ></div>

    <!-- Floating Elements -->
    <div class="absolute top-20 left-10 w-20 h-20 bg-white/10 rounded-full animate-float"></div>
    <div
      class="absolute top-40 right-20 w-16 h-16 bg-accent/20 rounded-full animate-float"
      style="animation-delay: -2s"
    ></div>
    <div
      class="absolute bottom-40 left-20 w-12 h-12 bg-secondary/20 rounded-full animate-float"
      style="animation-delay: -4s"
    ></div>

    <div class="relative z-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8 text-center">
      <div class="animate-slide-up">
        <div
          class="inline-flex items-center px-4 py-2 rounded-full bg-white/20 border border-white/30 backdrop-blur-sm mb-8"
        >
          <span class="w-2 h-2 bg-green-400 rounded-full mr-2 animate-pulse"></span>
          <span class="text-white font-medium">✨ Now with Advanced AI Categorization</span>
        </div>

        <h1 class="text-6xl md:text-8xl font-black text-white mb-6 leading-tight">
          Finally,
          <span
            class="gradient-text bg-clip-text text-transparent bg-gradient-to-r from-yellow-400 via-pink-500 to-cyan-400"
          >
            Inbox Zero
          </span>
          <br />is Within Reach
        </h1>

        <p class="text-xl md:text-2xl text-white/90 mb-12 max-w-3xl mx-auto leading-relaxed">
          Let our cutting-edge AI clean up your Gmail automatically, so you can focus on what truly
          matters.
          <span class="text-yellow-300 font-semibold">Join 50,000+ professionals</span> who've
          reclaimed their productivity.
        </p>

        <div class="flex flex-col sm:flex-row gap-4 justify-center items-center mb-16">
          <button
            class="group px-10 py-4 bg-white text-primary text-lg font-bold rounded-2xl shadow-2xl hover:shadow-white/25 hover:scale-105 transition-all duration-300 flex items-center space-x-2"
          >
            <span>Start Free Trial</span>
            <svg
              class="w-5 h-5 group-hover:translate-x-1 transition-transform"
              fill="none"
              stroke="currentColor"
              viewBox="0 0 24 24"
            >
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M13 7l5 5m0 0l-5 5m5-5H6"
              />
            </svg>
          </button>
          <button
            class="px-8 py-4 text-white border-2 border-white/30 rounded-2xl hover:bg-white/10 transition-all duration-300 backdrop-blur-sm"
          >
            Watch Demo (2 min)
          </button>
        </div>

        <!-- Stats -->
        <div class="grid grid-cols-3 gap-8 max-w-2xl mx-auto">
          <div class="text-center">
            <div class="text-3xl font-bold text-white mb-1">2.3M+</div>
            <div class="text-white/70">Emails Cleaned</div>
          </div>
          <div class="text-center">
            <div class="text-3xl font-bold text-white mb-1">15TB</div>
            <div class="text-white/70">Storage Freed</div>
          </div>
          <div class="text-center">
            <div class="text-3xl font-bold text-white mb-1">4.9/5</div>
            <div class="text-white/70">User Rating</div>
          </div>
        </div>
      </div>
    </div>

    <!-- Scroll Indicator -->
    <div class="absolute bottom-8 left-1/2 transform -translate-x-1/2 animate-bounce-slow">
      <div class="w-6 h-10 border-2 border-white/50 rounded-full flex justify-center">
        <div class="w-1 h-3 bg-white/70 rounded-full mt-2 animate-pulse"></div>
      </div>
    </div>
  </section>
</template>

<script setup></script>

<style scoped>
/* No scoped styles needed, using Tailwind CSS */
</style>
