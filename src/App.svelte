<script>
  import Header from './lib/components/Header.svelte';
  import Footer from './lib/components/Footer.svelte';
  import Home from './pages/Home.svelte';
</script>

<div class="min-h-screen flex flex-col">
  <Header />
  <main class="flex-1">
    <Home />
  </main>
  <Footer />
</div>
