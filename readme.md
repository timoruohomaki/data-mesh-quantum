# Data Mesh Quantum

### Ok, so what is this?

While the Data Mesh and data product obviously follows the concept defined by Zhamak Dehghani in her book [Data Mesh](https://www.oreilly.com/library/view/data-mesh/9781492092384/) and blog posts [1](https://martinfowler.com/articles/data-monolith-to-mesh.html) and [2](https://martinfowler.com/articles/data-mesh-principles.html), some of the topics trace back to the earlier work of Neal Ford [Building Evolutionary Architectures](https://oreil.ly/Lrd6t) and Eric Evans [Domain-Driven Design](https://www.oreilly.com/library/view/domain-driven-design-tackling/0321125215/).

The following definitions might be useful:

**Architecture Quantum**
As defined by Neal Ford, an architectural quantum is the smallest unit of architecture that can be independently deployed, has high functional cohesion, and includes all the "structural elements required for its function".

In the Data Mesh, the data product is an architectural quantum, the smallest unit that can be independently deployed and managed.

**Data Mesh**
Data Mesh can be defined by its four principles:
1. Domain-oriented ownership
2. Data as a product
3. Self-serve data platform
4. Federated computational governance

**Data Product**
Data as a product has the following characteristics:
+ _Discoverable_
+ _Addressable_
+ _Understandable_
+ _Trustworthy and truthful_
+ _Natively accessible_
+ _Interoperable and composable_
+ _Valuable on its own_
+ _Secure_

As defined by Dehghani: _A data product provides a set of explicitly defined and easy to use data sharing contracts. Each data product is autonomous and its life cycle and model are managed independently of others._

**Domain**
_A sphere of knowledge, influence, or activity._

**Fitness Function**
_An architectural fitness function provides an objective integrity assessment of some architectural characteristic(s)._

### Initial Plan

The following diagram illustrates the basic concept of Data Quantum.

![data-mesh-quantum drawio](https://github.com/timoruohomaki/data-mesh-quantum/assets/25956101/5d2a9a44-8f11-4673-b5cf-dcbbbd4c6694)

So, the plan is to implement here:
- [ ] Data dictionary to act as a service registry for discovery (based on Serf)
- [ ] The observability plane (likely using Sentry.io)
- [ ] The control plane (tbd, maybe Terraform)
- [ ] The data onboarding, transform and fitness functions
- [ ] The publishing point for interoperable data products
