<script>
import { page } from "$app/stores";
</script>

<div class="container mx-auto">
  <p>it be me, {$page.data.user.name}!</p>
  <p>Click <a class="link link-primary" href="/">here</a> to go home</p>
</div>
