package controller

type LoadBalancers []*LoadBalancer

func (l LoadBalancers) find(lb *LoadBalancer) int {
	for i, lbi := range l {
		if *lb.id == *lbi.id {
			return i
		}
	}
	return -1
}

func (l LoadBalancers) SyncState() LoadBalancers {
	var loadbalancers LoadBalancers
	for _, loadbalancer := range l {
		lb := loadbalancer.SyncState()
		loadbalancer.ResourceRecordSet.SyncState()
		loadbalancer.TargetGroups.SyncState()
		loadbalancer.Listeners.SyncState()
		if lb != nil {
			loadbalancers = append(loadbalancers, lb)
		}
	}
	return loadbalancers
}

func (l LoadBalancers) StripDesiredState() {
	for _, lb := range l {
		lb.DesiredLoadBalancer = nil
		if lb.ResourceRecordSet != nil {
			lb.ResourceRecordSet.DesiredResourceRecordSet = nil
		}
	}
}