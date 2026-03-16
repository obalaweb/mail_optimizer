<template>
  <section id="testimonials" class="py-24 bg-gradient-to-br from-primary/5 to-secondary/5">
    <div class="max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-20">
        <div
          class="inline-block px-4 py-2 bg-secondary/10 rounded-full text-secondary font-semibold mb-4"
        >
          💬 Customer Love
        </div>
        <h2 class="text-5xl md:text-6xl font-black text-slate-900 mb-6">What Our Users Say</h2>
        <p class="text-xl text-slate-600 max-w-3xl mx-auto">
          Join thousands of professionals who've transformed their email experience.
        </p>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-8">
        <!-- Testimonial 1 -->
        <div class="pricing-card p-8 rounded-3xl shadow-xl hover-lift">
          <div class="flex items-center mb-6">
            <div
              class="w-12 h-12 bg-gradient-to-r from-primary to-secondary rounded-full flex items-center justify-center text-white font-bold text-lg mr-4"
            >
              S
            </div>
            <div>
              <div class="font-bold text-slate-900">Sarah Chen</div>
              <div class="text-slate-600 text-sm">Product Manager at TechCorp</div>
            </div>
          </div>
          <div class="flex mb-4">
            <span class="text-yellow-400">★★★★★</span>
          </div>
          <p class="text-slate-700 text-lg leading-relaxed mb-4">
            "Gmail Optimizer completely transformed my workflow. I went from 15,000 unread emails to
            inbox zero in under 10 minutes. The AI categorization is incredibly accurate!"
          </p>
          <div class="text-primary font-semibold">Verified Purchase</div>
        </div>

        <!-- Testimonial 2 -->
        <div class="pricing-card p-8 rounded-3xl shadow-xl hover-lift">
          <div class="flex items-center mb-6">
            <div
              class="w-12 h-12 bg-gradient-to-r from-secondary to-accent rounded-full flex items-center justify-center text-white font-bold text-lg mr-4"
            >
              M
            </div>
            <div>
              <div class="font-bold text-slate-900">Marcus Williams</div>
              <div class="text-slate-600 text-sm">CEO at StartupXYZ</div>
            </div>
          </div>
          <div class="flex mb-4">
            <span class="text-yellow-400">★★★★★</span>
          </div>
          <p class="text-slate-700 text-lg leading-relaxed mb-4">
            "As a founder, my inbox was chaos. Gmail Optimizer's smart rules saved me 2 hours daily.
            The ROI is incredible—this tool pays for itself in the first week."
          </p>
          <div class="text-secondary font-semibold">Verified Purchase</div>
        </div>

        <!-- Testimonial 3 -->
        <div class="pricing-card p-8 rounded-3xl shadow-xl hover-lift">
          <div class="flex items-center mb-6">
            <div
              class="w-12 h-12 bg-gradient-to-r from-accent to-primary rounded-full flex items-center justify-center text-white font-bold text-lg mr-4"
            >
              A
            </div>
            <div>
              <div class="font-bold text-slate-900">Amanda Rodriguez</div>
              <div class="text-slate-600 text-sm">Marketing Director</div>
            </div>
          </div>
          <div class="flex mb-4">
            <span class="text-yellow-400">★★★★★</span>
          </div>
          <p class="text-slate-700 text-lg leading-relaxed mb-4">
            "The analytics dashboard is a game-changer. I can see exactly how much time I'm saving
            and which email types consume most of my attention. Pure genius!"
          </p>
          <div class="text-accent font-semibold">Verified Purchase</div>
        </div>

        <!-- Testimonial 4 -->
        <div class="pricing-card p-8 rounded-3xl shadow-xl hover-lift md:col-span-2 lg:col-span-1">
          <div class="flex items-center mb-6">
            <div
              class="w-12 h-12 bg-gradient-to-r from-pink-500 to-rose-500 rounded-full flex items-center justify-center text-white font-bold text-lg mr-4"
            >
              D
            </div>
            <div>
              <div class="font-bold text-slate-900">David Park</div>
              <div class="text-slate-600 text-sm">Freelance Designer</div>
            </div>
          </div>
          <div class="flex mb-4">
            <span class="text-yellow-400">★★★★★</span>
          </div>
          <p class="text-slate-700 text-lg leading-relaxed mb-4">
            "I was skeptical about AI handling my emails, but the security and precision blew me
            away. It's like having a personal assistant that never sleeps."
          </p>
          <div class="text-pink-500 font-semibold">Verified Purchase</div>
        </div>

        <!-- Testimonial 5 -->
        <div class="pricing-card p-8 rounded-3xl shadow-xl hover-lift md:col-span-2 lg:col-span-1">
          <div class="flex items-center mb-6">
            <div
              class="w-12 h-12 bg-gradient-to-r from-green-500 to-emerald-500 rounded-full flex items-center justify-center text-white font-bold text-lg mr-4"
            >
              L
            </div>
            <div>
              <div class="font-bold text-slate-900">Lisa Thompson</div>
              <div class="text-slate-600 text-sm">Operations Manager</div>
            </div>
          </div>
          <div class="flex mb-4">
            <span class="text-yellow-400">★★★★★</span>
          </div>
          <p class="text-slate-700 text-lg leading-relaxed mb-4">
            "The one-click cleanup feature recovered 8GB of storage instantly. My Gmail runs faster
            than ever, and I finally have peace of mind with my inbox."
          </p>
          <div class="text-green-500 font-semibold">Verified Purchase</div>
        </div>

        <!-- Testimonial 6 -->
        <div class="pricing-card p-8 rounded-3xl shadow-xl hover-lift lg:col-span-1">
          <div class="flex items-center mb-6">
            <div
              class="w-12 h-12 bg-gradient-to-r from-indigo-500 to-purple-500 rounded-full flex items-center justify-center text-white font-bold text-lg mr-4"
            >
              R
            </div>
            <div>
              <div class="font-bold text-slate-900">Robert Kim</div>
              <div class="text-slate-600 text-sm">Software Engineer</div>
            </div>
          </div>
          <div class="flex mb-4">
            <span class="text-yellow-400">★★★★★</span>
          </div>
          <p class="text-slate-700 text-lg leading-relaxed mb-4">
            "The custom rules engine is incredibly powerful. I've automated complex workflows that
            would have taken hours to set up manually. Highly recommend!"
          </p>
          <div class="text-indigo-500 font-semibold">Verified Purchase</div>
        </div>
      </div>

      <!-- Trust indicators -->
      <div class="grid grid-cols-2 md:grid-cols-4 gap-8 mt-16 text-center">
        <div>
          <div class="text-3xl font-bold text-slate-900 mb-2">50K+</div>
          <div class="text-slate-600">Happy Users</div>
        </div>
        <div>
          <div class="text-3xl font-bold text-slate-900 mb-2">99.9%</div>
          <div class="text-slate-600">Uptime</div>
        </div>
        <div>
          <div class="text-3xl font-bold text-slate-900 mb-2">4.9/5</div>
          <div class="text-slate-600">App Store Rating</div>
        </div>
        <div>
          <div class="text-3xl font-bold text-slate-900 mb-2">24/7</div>
          <div class="text-slate-600">Support</div>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup></script>

<style scoped>
/* Tailwind CSS is used */
</style>
