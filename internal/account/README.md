# Account context
This bounded context is responsible for account management. This context is quite coarse-grained, as its responsibilities
are the following:
- Account registration
- Authentication
- Settings & profile settings (e.g. changing look, motto - you name it)
- Account deletion

There might be a few things not mentioned here - this README won't be frequently updated. 

There are some database migrations for this bounded context. It's included in the makefile.

