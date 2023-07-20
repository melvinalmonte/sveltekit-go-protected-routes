<div class="container mx-auto">
	<div class="navbar bg-base-100">
		<div class="flex-1">
			<p class="normal-case text-xl">Protected Routes Practice</p>
		</div>
	</div>
	<form method="POST">
		<div class="form-control w-full max-w-xs pt-4 pb-4">
			<label for="username" class="label">
				<span class="label-text">Username:</span>
			</label>
			<input
				class="input input-bordered input-primary w-full max-w-xs"
				id="username"
				name="username"
				type="text"
				placeholder="ie: Mick Mouse"
				autocomplete="off"
			/>
		</div>
		<div class="form-control w-full max-w-xs pt-4 pb-4">
			<label for="password" class="label">
				<span class="label-text">Password:</span>
			</label>
			<input
				class="input input-bordered input-primary w-full max-w-xs"
				id="password"
				name="password"
				type="password"
				placeholder="********"
				autocomplete="off"
			/>
		</div>

		<button class="btn btn-primary" type="submit">Log in</button>
	</form>
</div>
