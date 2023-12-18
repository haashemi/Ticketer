<main class="flex min-h-screen items-center justify-center">
	<div class="flex w-80 flex-col items-center justify-center gap-6 rounded-xl border border-slate-700 p-7">
		<h1 class="text-2xl font-semibold">Sign Up</h1>

		<form class="flex w-full flex-col gap-3">
			<label class="form-control w-full max-w-xs">
				<div class="label">
					<span class="label-text">Full Name</span>
				</div>
				<input type="password" placeholder="Reza Boghzi" class="input input-bordered w-full max-w-xs" />
			</label>

			<label class="form-control w-full max-w-xs">
				<div class="label">
					<span class="label-text">EMail</span>
				</div>
				<input type="email" placeholder="username@gmail.com" class="input input-bordered w-full max-w-xs" />
			</label>

			<label class="form-control w-full max-w-xs">
				<div class="label">
					<span class="label-text">Password</span>
				</div>
				<input type="password" placeholder="********" class="input input-bordered w-full max-w-xs" />
			</label>

			<div class="flex justify-end">
				<button class="btn btn-primary">Sign Up</button>
			</div>
		</form>

		<p>Already have an account? <a class="text-secondary" href="/signin">Sign In!</a></p>
	</div>
</main>
