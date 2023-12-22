<div class="absolute left-1/2 top-1/2 -z-10 h-96 w-96 -translate-x-1/2 -translate-y-1/2 rounded-full bg-blue-900 blur-3xl" />

<main class="flex min-h-screen items-center justify-center bg-black/80 backdrop-blur-xl">
	<div class="flex flex-col items-center justify-center gap-6 rounded-xl border border-slate-700 p-10">
		<h1 class="text-2xl font-semibold">Sign Up</h1>

		<form class="flex w-full flex-col gap-3">
			<label class="form-control w-full max-w-xs">
				<div class="label">
					<span class="label-text">Full Name</span>
				</div>
				<input required type="text" placeholder="Reza Boghzi" class="input input-bordered w-full max-w-xs" />
			</label>

			<label class="form-control w-full max-w-xs">
				<div class="label">
					<span class="label-text">Email</span>
				</div>
				<input required type="email" placeholder="username@gmail.com" class="input input-bordered w-full max-w-xs" />
			</label>

			<label class="form-control w-full max-w-xs">
				<div class="label">
					<span class="label-text">Password</span>
				</div>
				<input required type="password" placeholder="********" class="input input-bordered w-full max-w-xs" />
			</label>

			<div class="flex justify-end">
				<button class="btn btn-primary">Sign Up</button>
			</div>
		</form>

		<p>Already have an account? <a class="text-primary" href="/signin">Sign In!</a></p>
	</div>
</main>
