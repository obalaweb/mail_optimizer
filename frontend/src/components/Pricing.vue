<template>
  <section
    id="pricing"
    class="py-24 bg-gradient-to-br from-slate-900 to-primary relative overflow-hidden"
  >
    <!-- Background Effects -->
    <div class="absolute inset-0 bg-gradient-to-r from-primary/20 to-secondary/20"></div>
    <div
      class="absolute top-0 left-0 w-72 h-72 bg-accent/10 rounded-full -translate-x-1/2 -translate-y-1/2 animate-pulse-slow"
    ></div>
    <div
      class="absolute bottom-0 right-0 w-96 h-96 bg-secondary/10 rounded-full translate-x-1/3 translate-y-1/3 animate-pulse-slow"
      style="animation-delay: -2s"
    ></div>

    <div class="relative z-10 max-w-7xl mx-auto px-4 sm:px-6 lg:px-8">
      <div class="text-center mb-20">
        <div class="inline-block px-4 py-2 bg-white/10 rounded-full text-white font-semibold mb-4">
          💎 Simple Pricing
        </div>
        <h2 class="text-5xl md:text-6xl font-black text-white mb-6">
          Choose Your
          <span
            class="gradient-text bg-clip-text text-transparent bg-gradient-to-r from-yellow-400 to-cyan-400"
          >
            Perfect Plan
          </span>
        </h2>
        <p class="text-xl text-white/80 max-w-3xl mx-auto mb-8">
          Start free, upgrade when you're ready. No hidden fees, cancel anytime.
        </p>

        <!-- Billing Toggle -->
        <div class="flex items-center justify-center space-x-4 mb-12">
          <span class="text-white/70">Monthly</span>
          <button
            class="relative inline-flex items-center w-16 h-8 bg-white/20 rounded-full transition-colors duration-300"
          >
            <span
              class="absolute left-1 w-6 h-6 bg-white rounded-full shadow-lg transform transition-transform duration-300 translate-x-8"
            ></span>
          </button>
          <span class="text-white font-semibold">
            Annual
            <span class="inline-block px-2 py-1 bg-green-500 text-white text-xs rounded-full ml-1"
              >Save 20%</span
            >
          </span>
        </div>
      </div>

      <div class="grid grid-cols-1 md:grid-cols-3 gap-8 max-w-6xl mx-auto">
        <!-- Free Plan -->
        <div class="pricing-card p-8 rounded-3xl shadow-2xl hover-lift">
          <div class="text-center mb-8">
            <div
              class="inline-block px-4 py-2 bg-slate-100 rounded-full text-slate-700 font-semibold mb-4"
            >
              Perfect for Starters
            </div>
            <h3 class="text-3xl font-bold text-slate-900 mb-2">Free</h3>
            <div class="text-5xl font-black text-slate-900 mb-2">
              $0
              <span class="text-xl text-slate-600 font-normal">/month</span>
            </div>
            <p class="text-slate-600">Get started with basic cleanup</p>
          </div>

          <ul class="space-y-4 mb-8">
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700">100 email actions/month</span>
            </li>
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700">Basic auto-cleanup</span>
            </li>
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700">Email support</span>
            </li>
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700">Basic analytics</span>
            </li>
          </ul>

          <button
            class="w-full px-6 py-4 bg-slate-900 text-white rounded-2xl hover:bg-slate-800 transition-all duration-300 font-semibold"
          >
            Get Started Free
          </button>
        </div>

        <!-- Pro Plan -->
        <div
          class="pricing-card p-8 rounded-3xl shadow-2xl hover-lift relative border-4 border-primary/50"
        >
          <!-- Popular badge -->
          <div class="absolute -top-4 left-1/2 transform -translate-x-1/2">
            <div
              class="px-6 py-2 bg-gradient-to-r from-primary to-secondary text-white rounded-full font-bold text-sm shadow-lg"
            >
              🔥 Most Popular
            </div>
          </div>

          <div class="text-center mb-8">
            <div
              class="inline-block px-4 py-2 bg-primary/10 rounded-full text-primary font-semibold mb-4"
            >
              For Power Users
            </div>
            <h3 class="text-3xl font-bold text-slate-900 mb-2">Pro</h3>
            <div class="text-5xl font-black text-slate-900 mb-2">
              $9
              <span class="text-xl text-slate-600 font-normal">/month</span>
            </div>
            <p class="text-slate-600">Everything you need to stay organized</p>
          </div>

          <ul class="space-y-4 mb-8">
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700"><strong>5,000</strong> email actions/month</span>
            </li>
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700">Advanced auto-cleanup</span>
            </li>
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700"><strong>AI categorization</strong></span>
            </li>
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700">Custom rules engine</span>
            </li>
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700">Advanced analytics</span>
            </li>
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700">Priority support</span>
            </li>
          </ul>

          <button
            class="w-full px-6 py-4 bg-gradient-to-r from-primary to-secondary text-white rounded-2xl hover:shadow-lg hover:scale-105 transition-all duration-300 font-semibold"
          >
            Start Pro Trial
          </button>
        </div>

        <!-- Team Plan -->
        <div class="pricing-card p-8 rounded-3xl shadow-2xl hover-lift">
          <div class="text-center mb-8">
            <div
              class="inline-block px-4 py-2 bg-secondary/10 rounded-full text-secondary font-semibold mb-4"
            >
              For Teams & Enterprises
            </div>
            <h3 class="text-3xl font-bold text-slate-900 mb-2">Team</h3>
            <div class="text-5xl font-black text-slate-900 mb-2">
              $29
              <span class="text-xl text-slate-600 font-normal">/month</span>
            </div>
            <p class="text-slate-600">Scale across your entire organization</p>
          </div>

          <ul class="space-y-4 mb-8">
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700"><strong>Unlimited</strong> email actions</span>
            </li>
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700">All Pro features</span>
            </li>
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700">Multi-user dashboard</span>
            </li>
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700">Team analytics</span>
            </li>
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700">SSO integration</span>
            </li>
            <li class="flex items-center">
              <svg
                class="w-5 h-5 text-green-500 mr-3"
                fill="none"
                stroke="currentColor"
                viewBox="0 0 24 24"
              >
                <path
                  stroke-linecap="round"
                  stroke-linejoin="round"
                  stroke-width="2"
                  d="M5 13l4 4L19 7"
                />
              </svg>
              <span class="text-slate-700">Dedicated success manager</span>
            </li>
          </ul>

          <button
            class="w-full px-6 py-4 bg-secondary text-white rounded-2xl hover:bg-secondary/90 transition-all duration-300 font-semibold"
          >
            Contact Sales
          </button>
        </div>
      </div>

      <!-- Money back guarantee -->
      <div class="text-center mt-16">
        <div class="inline-flex items-center px-6 py-3 bg-white/10 rounded-full text-white">
          <svg class="w-6 h-6 mr-2" fill="none" stroke="currentColor" viewBox="0 0 24 24">
            <path
              stroke-linecap="round"
              stroke-linejoin="round"
              stroke-width="2"
              d="M9 12l2 2 4-4m5.618-4.016A11.955 11.955 0 0112 2.944a11.955 11.955 0 01-8.618 3.04A12.02 12.02 0 003 9c0 5.591 3.824 10.29 9 11.622 5.176-1.332 9-6.03 9-11.622 0-1.042-.133-2.052-.382-3.016z"
            />
          </svg>
          <span class="font-semibold">30-day money-back guarantee • Cancel anytime</span>
        </div>
      </div>
    </div>
  </section>
</template>

<script setup></script>

<style scoped>
/* Tailwind CSS is used */
</style>
