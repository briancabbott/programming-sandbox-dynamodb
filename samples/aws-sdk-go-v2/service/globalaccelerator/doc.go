// Code generated by smithy-go-codegen DO NOT EDIT.

// Package globalaccelerator provides the API client, operations, and parameter
// types for AWS Global Accelerator.
//
// AWS Global Accelerator This is the AWS Global Accelerator API Reference. This
// guide is for developers who need detailed information about AWS Global
// Accelerator API actions, data types, and errors. For more information about
// Global Accelerator features, see the AWS Global Accelerator Developer Guide
// (https://docs.aws.amazon.com/global-accelerator/latest/dg/Welcome.html). AWS
// Global Accelerator is a service in which you create accelerators to improve
// availability and performance of your applications for local and global users.
// Global Accelerator directs traffic to optimal endpoints over the AWS global
// network. This improves the availability and performance of your internet
// applications that are used by a global audience. Global Accelerator is a global
// service that supports endpoints in multiple AWS Regions, which are listed in the
// AWS Region Table
// (https://aws.amazon.com/about-aws/global-infrastructure/regional-product-services/).
// Global Accelerator is a global service that supports endpoints in multiple AWS
// Regions but you must specify the US West (Oregon) Region to create or update
// accelerators. By default, Global Accelerator provides you with static IP
// addresses that you associate with your accelerator. (Instead of using the IP
// addresses that Global Accelerator provides, you can configure these entry points
// to be IPv4 addresses from your own IP address ranges that you bring to Global
// Accelerator.) The static IP addresses are anycast from the AWS edge network and
// distribute incoming application traffic across multiple endpoint resources in
// multiple AWS Regions, which increases the availability of your applications.
// Endpoints can be Network Load Balancers, Application Load Balancers, EC2
// instances, or Elastic IP addresses that are located in one AWS Region or
// multiple Regions. Global Accelerator uses the AWS global network to route
// traffic to the optimal regional endpoint based on health, client location, and
// policies that you configure. The service reacts instantly to changes in health
// or configuration to ensure that internet traffic from clients is directed to
// only healthy endpoints. Global Accelerator includes components that work
// together to help you improve performance and availability for your applications:
// Static IP address By default, AWS Global Accelerator provides you with a set of
// static IP addresses that are anycast from the AWS edge network and serve as the
// single fixed entry points for your clients. Or you can configure these entry
// points to be IPv4 addresses from your own IP address ranges that you bring to
// Global Accelerator (BYOIP). For more information, see Bring Your Own IP
// Addresses (BYOIP)
// (https://docs.aws.amazon.com/global-accelerator/latest/dg/using-byoip.html) in
// the AWS Global Accelerator Developer Guide. If you already have load balancers,
// EC2 instances, or Elastic IP addresses set up for your applications, you can
// easily add those to Global Accelerator to allow the resources to be accessed by
// the static IP addresses. The static IP addresses remain assigned to your
// accelerator for as long as it exists, even if you disable the accelerator and it
// no longer accepts or routes traffic. However, when you delete an accelerator,
// you lose the static IP addresses that are assigned to it, so you can no longer
// route traffic by using them. You can use IAM policies with Global Accelerator to
// limit the users who have permissions to delete an accelerator. For more
// information, see Authentication and Access Control
// (https://docs.aws.amazon.com/global-accelerator/latest/dg/auth-and-access-control.html)
// in the AWS Global Accelerator Developer Guide. Accelerator An accelerator
// directs traffic to optimal endpoints over the AWS global network to improve
// availability and performance for your internet applications that have a global
// audience. Each accelerator includes one or more listeners. DNS name Global
// Accelerator assigns each accelerator a default Domain Name System (DNS) name,
// similar to a1234567890abcdef.awsglobalaccelerator.com, that points to your
// Global Accelerator static IP addresses. Depending on the use case, you can use
// your accelerator's static IP addresses or DNS name to route traffic to your
// accelerator, or set up DNS records to route traffic using your own custom domain
// name. Network zone A network zone services the static IP addresses for your
// accelerator from a unique IP subnet. Similar to an AWS Availability Zone, a
// network zone is an isolated unit with its own set of physical infrastructure.
// When you configure an accelerator, by default, Global Accelerator allocates two
// IPv4 addresses for it. If one IP address from a network zone becomes unavailable
// due to IP address blocking by certain client networks, or network disruptions,
// then client applications can retry on the healthy static IP address from the
// other isolated network zone. Listener A listener processes inbound connections
// from clients to Global Accelerator, based on the protocol and port that you
// configure. Each listener has one or more endpoint groups associated with it, and
// traffic is forwarded to endpoints in one of the groups. You associate endpoint
// groups with listeners by specifying the Regions that you want to distribute
// traffic to. Traffic is distributed to optimal endpoints within the endpoint
// groups associated with a listener. Endpoint group Each endpoint group is
// associated with a specific AWS Region. Endpoint groups include one or more
// endpoints in the Region. You can increase or reduce the percentage of traffic
// that would be otherwise directed to an endpoint group by adjusting a setting
// called a traffic dial. The traffic dial lets you easily do performance testing
// or blue/green deployment testing for new releases across different AWS Regions,
// for example. Endpoint An endpoint is a Network Load Balancer, Application Load
// Balancer, EC2 instance, or Elastic IP address. Traffic is routed to endpoints
// based on several factors, including the geo-proximity to the user, the health of
// the endpoint, and the configuration options that you choose, such as endpoint
// weights. For each endpoint, you can configure weights, which are numbers that
// you can use to specify the proportion of traffic to route to each one. This can
// be useful, for example, to do performance testing within a Region.
package globalaccelerator
