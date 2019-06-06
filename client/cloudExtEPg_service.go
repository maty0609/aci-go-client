package client

import (
	"fmt"

	"github.com/ciscoecosystem/aci-go-client/models"
	"github.com/ciscoecosystem/aci-go-client/container"
	"github.com/hashicorp/terraform/helper/schema"
	



	


)









func (sm *ServiceManager) CreateCloudExternalEPg(name string ,cloud_application_container string ,tenant string , description string, cloudExtEPgattr models.CloudExternalEPgAttributes) (*models.CloudExternalEPg, error) {	
	rn := fmt.Sprintf("cloudextepg-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/cloudapp-%s", tenant ,cloud_application_container )
	cloudExtEPg := models.NewCloudExternalEPg(rn, parentDn, description, cloudExtEPgattr)
	err := sm.Save(cloudExtEPg)
	return cloudExtEPg, err
}

func (sm *ServiceManager) ReadCloudExternalEPg(name string ,cloud_application_container string ,tenant string ) (*models.CloudExternalEPg, error) {
	dn := fmt.Sprintf("uni/tn-%s/cloudapp-%s/cloudextepg-%s", tenant ,cloud_application_container ,name )    
	cont, err := sm.Get(dn)
	if err != nil {
		return nil, err
	}

	cloudExtEPg := models.CloudExternalEPgFromContainer(cont)
	return cloudExtEPg, nil
}

func (sm *ServiceManager) DeleteCloudExternalEPg(name string ,cloud_application_container string ,tenant string ) error {
	dn := fmt.Sprintf("uni/tn-%s/cloudapp-%s/cloudextepg-%s", tenant ,cloud_application_container ,name )
	return sm.DeleteByDn(dn, models.CloudextepgClassName)
}

func (sm *ServiceManager) UpdateCloudExternalEPg(name string ,cloud_application_container string ,tenant string  ,description string, cloudExtEPgattr models.CloudExternalEPgAttributes) (*models.CloudExternalEPg, error) {
	rn := fmt.Sprintf("cloudextepg-%s",name)
	parentDn := fmt.Sprintf("uni/tn-%s/cloudapp-%s", tenant ,cloud_application_container )
	cloudExtEPg := models.NewCloudExternalEPg(rn, parentDn, description, cloudExtEPgattr)

    cloudExtEPg.Status = "modified"
	err := sm.Save(cloudExtEPg)
	return cloudExtEPg, err

}

func (sm *ServiceManager) ListCloudExternalEPg(cloud_application_container string ,tenant string ) ([]*models.CloudExternalEPg, error) {

	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/uni/tn-%s/cloudapp-%s/cloudExtEPg.json", baseurlStr , tenant ,cloud_application_container )
    
    cont, err := sm.GetViaURL(dnUrl)
	list := models.CloudExternalEPgListFromContainer(cont)

	return list, err
}

func (sm *ServiceManager) CreateRelationfvRsSecInheritedFromCloudExternalEPg( parentDn, tDn string) error {
	dn := fmt.Sprintf("%s/rssecInherited-[%s]", parentDn, tDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsSecInherited", dn))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) DeleteRelationfvRsSecInheritedFromCloudExternalEPg(parentDn , tDn string) error{
	dn := fmt.Sprintf("%s/rssecInherited-[%s]", parentDn, tDn)
	return sm.DeleteByDn(dn , "fvRsSecInherited")
}

func (sm *ServiceManager) ReadRelationfvRsSecInheritedFromCloudExternalEPg( parentDn string) (interface{},error) {
	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/%s/%s.json",baseurlStr,parentDn,"fvRsSecInherited")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont,"fvRsSecInherited")
	
	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList{
		dat := models.G(contItem, "tDn")
		st.Add(dat)
	}
	return st, err
			





}
func (sm *ServiceManager) CreateRelationfvRsProvFromCloudExternalEPg( parentDn, tnVzBrCPName string) error {
	dn := fmt.Sprintf("%s/rsprov-%s", parentDn, tnVzBrCPName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsProv", dn))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) DeleteRelationfvRsProvFromCloudExternalEPg(parentDn , tnVzBrCPName string) error{
	dn := fmt.Sprintf("%s/rsprov-%s", parentDn, tnVzBrCPName)
	return sm.DeleteByDn(dn , "fvRsProv")
}

func (sm *ServiceManager) ReadRelationfvRsProvFromCloudExternalEPg( parentDn string) (interface{},error) {
	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/%s/%s.json",baseurlStr,parentDn,"fvRsProv")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont,"fvRsProv")
	
	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList{
		dat := models.G(contItem, "tnVzBrCPName")
		st.Add(dat)
	}
	return st, err
			





}
func (sm *ServiceManager) CreateRelationfvRsConsIfFromCloudExternalEPg( parentDn, tnVzCPIfName string) error {
	dn := fmt.Sprintf("%s/rsconsIf-%s", parentDn, tnVzCPIfName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsConsIf", dn))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) DeleteRelationfvRsConsIfFromCloudExternalEPg(parentDn , tnVzCPIfName string) error{
	dn := fmt.Sprintf("%s/rsconsIf-%s", parentDn, tnVzCPIfName)
	return sm.DeleteByDn(dn , "fvRsConsIf")
}

func (sm *ServiceManager) ReadRelationfvRsConsIfFromCloudExternalEPg( parentDn string) (interface{},error) {
	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/%s/%s.json",baseurlStr,parentDn,"fvRsConsIf")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont,"fvRsConsIf")
	
	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList{
		dat := models.G(contItem, "tnVzCPIfName")
		st.Add(dat)
	}
	return st, err
			





}
func (sm *ServiceManager) CreateRelationfvRsCustQosPolFromCloudExternalEPg( parentDn, tnQosCustomPolName string) error {
	dn := fmt.Sprintf("%s/rscustQosPol", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnQosCustomPolName": "%s"
								
			}
		}
	}`, "fvRsCustQosPol", dn,tnQosCustomPolName))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) ReadRelationfvRsCustQosPolFromCloudExternalEPg( parentDn string) (interface{},error) {
	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/%s/%s.json",baseurlStr,parentDn,"fvRsCustQosPol")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont,"fvRsCustQosPol")
	
	if len(contList) > 0 {
		dat := models.G(contList[0], "tnQosCustomPolName")
		return dat, err
	} else {
		return nil,err
	}
		





}
func (sm *ServiceManager) CreateRelationfvRsConsFromCloudExternalEPg( parentDn, tnVzBrCPName string) error {
	dn := fmt.Sprintf("%s/rscons-%s", parentDn, tnVzBrCPName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsCons", dn))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) DeleteRelationfvRsConsFromCloudExternalEPg(parentDn , tnVzBrCPName string) error{
	dn := fmt.Sprintf("%s/rscons-%s", parentDn, tnVzBrCPName)
	return sm.DeleteByDn(dn , "fvRsCons")
}

func (sm *ServiceManager) ReadRelationfvRsConsFromCloudExternalEPg( parentDn string) (interface{},error) {
	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/%s/%s.json",baseurlStr,parentDn,"fvRsCons")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont,"fvRsCons")
	
	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList{
		dat := models.G(contItem, "tnVzBrCPName")
		st.Add(dat)
	}
	return st, err
			





}
func (sm *ServiceManager) CreateRelationcloudRsCloudEPgCtxFromCloudExternalEPg( parentDn, tnFvCtxName string) error {
	dn := fmt.Sprintf("%s/rsCloudEPgCtx", parentDn)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s","tnFvCtxName": "%s"
								
			}
		}
	}`, "cloudRsCloudEPgCtx", dn,tnFvCtxName))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) ReadRelationcloudRsCloudEPgCtxFromCloudExternalEPg( parentDn string) (interface{},error) {
	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/%s/%s.json",baseurlStr,parentDn,"cloudRsCloudEPgCtx")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont,"cloudRsCloudEPgCtx")
	
	if len(contList) > 0 {
		dat := models.G(contList[0], "tnFvCtxName")
		return dat, err
	} else {
		return nil,err
	}
		





}
func (sm *ServiceManager) CreateRelationfvRsProtByFromCloudExternalEPg( parentDn, tnVzTabooName string) error {
	dn := fmt.Sprintf("%s/rsprotBy-%s", parentDn, tnVzTabooName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsProtBy", dn))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) DeleteRelationfvRsProtByFromCloudExternalEPg(parentDn , tnVzTabooName string) error{
	dn := fmt.Sprintf("%s/rsprotBy-%s", parentDn, tnVzTabooName)
	return sm.DeleteByDn(dn , "fvRsProtBy")
}

func (sm *ServiceManager) ReadRelationfvRsProtByFromCloudExternalEPg( parentDn string) (interface{},error) {
	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/%s/%s.json",baseurlStr,parentDn,"fvRsProtBy")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont,"fvRsProtBy")
	
	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList{
		dat := models.G(contItem, "tnVzTabooName")
		st.Add(dat)
	}
	return st, err
			





}
func (sm *ServiceManager) CreateRelationfvRsIntraEpgFromCloudExternalEPg( parentDn, tnVzBrCPName string) error {
	dn := fmt.Sprintf("%s/rsintraEpg-%s", parentDn, tnVzBrCPName)
	containerJSON := []byte(fmt.Sprintf(`{
		"%s": {
			"attributes": {
				"dn": "%s"				
			}
		}
	}`, "fvRsIntraEpg", dn))

	jsonPayload, err := container.ParseJSON(containerJSON)
	if err != nil {
		return err
	}

	req, err := sm.client.MakeRestRequest("POST", fmt.Sprintf("%s.json", sm.MOURL), jsonPayload, true)
	if err != nil {
		return err
	}

	cont, _, err := sm.client.Do(req)
	if err != nil {
		return err
	}
	fmt.Printf("%+v", cont)

	return nil
}

func (sm *ServiceManager) DeleteRelationfvRsIntraEpgFromCloudExternalEPg(parentDn , tnVzBrCPName string) error{
	dn := fmt.Sprintf("%s/rsintraEpg-%s", parentDn, tnVzBrCPName)
	return sm.DeleteByDn(dn , "fvRsIntraEpg")
}

func (sm *ServiceManager) ReadRelationfvRsIntraEpgFromCloudExternalEPg( parentDn string) (interface{},error) {
	baseurlStr := "/api/node/class"	
	dnUrl := fmt.Sprintf("%s/%s/%s.json",baseurlStr,parentDn,"fvRsIntraEpg")
	cont, err := sm.GetViaURL(dnUrl)

	contList := models.ListFromContainer(cont,"fvRsIntraEpg")
	
	st := &schema.Set{
		F: schema.HashString,
	}
	for _, contItem := range contList{
		dat := models.G(contItem, "tnVzBrCPName")
		st.Add(dat)
	}
	return st, err
			





}

