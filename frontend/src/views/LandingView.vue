<template>
  <div class="flex flex-col min-h-screen bg-gray-100">
    <Header />
    <Hero />
    <Features />
    <SocialProof />
    <Testimonials />
    <Pricing />
    <Footer />
  </div>
</template>

<script setup>
import Header from '../components/Header.vue'
import Hero from '../components/Hero.vue'
import Features from '../components/Features.vue'
import Testimonials from '../components/Testimonials.vue'
import Pricing from '../components/Pricing.vue'
import Footer from '../components/Footer.vue'
import SocialProof from '@/components/SocialProof.vue'
</script>

<style scoped>
/* No scoped styles needed, using Tailwind CSS */
</style>
