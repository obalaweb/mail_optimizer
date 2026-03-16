<template>
  <nav class="fixed top-0 w-full z-50 glass-morphism animate-fade-in">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="flex justify-between items-center h-16">
        <div class="flex items-center space-x-3">
          <div
            class="w-10 h-10 bg-gradient-to-r from-primary to-secondary rounded-xl flex items-center justify-center animate-glow"
          >
            <svg class="w-6 h-6 text-white" fill="currentColor" viewBox="0 0 20 20">
              <path d="M2.003 5.884L10 9.882l7.997-3.998A2 2 0 0016 4H4a2 2 0 00-1.997 1.884z" />
              <path d="M18 8.118l-8 4-8-4V14a2 2 0 002 2h12a2 2 0 002-2V8.118z" />
            </svg>
          </div>
          <span class="text-2xl font-bold gradient-text">Gmail Optimizer</span>
        </div>
        <div class="hidden md:flex items-center space-x-8">
          <a
            href="#features"
            class="text-slate-700 hover:text-primary transition-colors font-medium hover:underline decoration-2 underline-offset-4"
            >Features</a
          >
          <a
            href="#pricing"
            class="text-slate-700 hover:text-primary transition-colors font-medium hover:underline decoration-2 underline-offset-4"
            >Pricing</a
          >
          <a
            href="#testimonials"
            class="text-slate-700 hover:text-primary transition-colors font-medium hover:underline decoration-2 underline-offset-4"
            >Reviews</a
          >
          <button
            class="px-6 py-2 bg-gradient-to-r from-primary to-secondary text-white rounded-full hover:shadow-lg hover:scale-105 transition-all duration-300 font-medium"
          >
            Get Started Free
          </button>
        </div>
        <div class="md:hidden">
          <button class="text-slate-700 p-2">
            <svg class="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M4 6h16M4 12h16M4 18h16"
              />
            </svg>
          </button>
        </div>
      </div>
    </div>
  </nav>
</template>

<script setup></script>

<style scoped>
/* No scoped styles needed, using Tailwind CSS */
</style>
