<script>
  import { Github } from 'lucide-svelte';
</script>

<header class="bg-slate-900 text-white">
  <nav class="max-w-6xl mx-auto px-4 py-4 flex justify-between items-center">
    <a href="/" class="text-xl font-bold text-white hover:opacity-90 transition-opacity">
      Bye Bye Server
    </a>
    <a
      href="https://github.com/realignedtech/bye-bye-server"
      target="_blank"
      rel="noopener noreferrer"
      class="flex items-center gap-2 text-white/90 text-sm font-medium transition-colors hover:text-white"
    >
      <Github class="w-5 h-5" />
      <span class="hidden sm:inline">GitHub</span>
    </a>
  </nav>
</header>
