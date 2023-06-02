package neo4j

import (
	"github.com/c12s/oort/domain/model"
	"github.com/c12s/oort/domain/store/acl"
)

type cypherFactory interface {
	createResourceCypher(req acl.CreateResourceReq) (string, map[string]interface{})
	deleteResourceCypher(req acl.DeleteResourceReq) (string, map[string]interface{})
	getResourceCypher(req acl.GetResourceReq) (string, map[string]interface{})
	createAttributeCypher(req acl.CreateAttributeReq) (string, map[string]interface{})
	updateAttributeCypher(req acl.UpdateAttributeReq) (string, map[string]interface{})
	deleteAttributeCypher(req acl.DeleteAttributeReq) (string, map[string]interface{})
	createAggregationRelCypher(req acl.CreateAggregationRelReq) (string, map[string]interface{})
	deleteAggregationRelCypher(req acl.DeleteAggregationRelReq) (string, map[string]interface{})
	createCompositionRelCypher(req acl.CreateCompositionRelReq) (string, map[string]interface{})
	deleteCompositionRelCypher(req acl.DeleteCompositionRelReq) (string, map[string]interface{})
	createPermissionCypher(req acl.CreatePermissionReq) (string, map[string]interface{})
	deletePermissionCypher(req acl.DeletePermissionReq) (string, map[string]interface{})
	getEffectivePermissionsWithPriorityCypher(req acl.GetPermissionHierarchyReq) (string, map[string]interface{})
}

type nonCachedPermissionsCypherFactory struct {
}

func NewNonCachedPermissionsCypherFactory() cypherFactory {
	return &nonCachedPermissionsCypherFactory{}
}

const ncCreateResourceQuery = `
MERGE (r:Resource{name: $name})
MERGE (root:Resource{name: $rootName})
MERGE (root)-[:Includes{kind: $composition}]->(r)
`

func (f nonCachedPermissionsCypherFactory) createResourceCypher(req acl.CreateResourceReq) (string, map[string]interface{}) {
	return ncCreateResourceQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"rootName":    model.RootResource.Name(),
			"composition": model.CompositionRelationship}
}

const ncDeleteResourceQuery = `
MATCH (r:Resource{name: $name})
WITH r
// nadji sve resurse za brisanje
CALL {
    WITH r
    OPTIONAL MATCH (r)-[:Includes*{kind: $composition}]->(d:Resource)
    RETURN collect(d) + collect(r) AS delRes
}
// obrisi sve atribute resursa za brisanje
CALL {
    WITH delRes
    UNWIND delRes AS r
    MATCH (r)-[:Includes{kind: $composition}]->(a:Attribute)
    DETACH DELETE a
}
// obrisi sve direktno dodeljene dozvole resursa za brisanje
CALL {
    WITH delRes
    UNWIND delRes AS r
    MATCH (r)-[:Has|On]-(p:Permission)
    DETACH DELETE p
}
// obrisi resurse
CALL {
    WITH delRes
    UNWIND delRes AS r
    DETACH DELETE r
}
`

func (f nonCachedPermissionsCypherFactory) deleteResourceCypher(req acl.DeleteResourceReq) (string, map[string]interface{}) {
	return ncDeleteResourceQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

const ncGetResourceQuery = `
MATCH (resource:Resource{name: $name})
OPTIONAL MATCH (attr:Attribute)<-[:Includes{kind: $composition}]-(resource)
RETURN resource.name, collect(properties(attr)) as attrs
`

func (f nonCachedPermissionsCypherFactory) getResourceCypher(req acl.GetResourceReq) (string, map[string]interface{}) {
	return ncGetResourceQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"composition": model.CompositionRelationship}
}

const ncCreateAttributeQuery = `
MATCH (r:Resource{name: $name})
WHERE NOT ((r)-[:Includes{kind: $composition}]->(:Attribute{name: $attrName}))
CREATE (r)-[:Includes{kind: $composition}]->(:Attribute{name: $attrName, kind: $attrKind, value: $attrValue})
`

func (f nonCachedPermissionsCypherFactory) createAttributeCypher(req acl.CreateAttributeReq) (string, map[string]interface{}) {
	return ncCreateAttributeQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"attrName":    req.Attribute.Name(),
			"attrKind":    req.Attribute.Kind(),
			"attrValue":   req.Attribute.Value(),
			"composition": model.CompositionRelationship}
}

const ncUpdateAttributeQuery = `
MATCH ((:Resource{name: $name})-[:Includes{kind: $composition}]->(a:Attribute{name: $attrName, kind: $attrKind}))
SET a.value = $attrValue
`

func (f nonCachedPermissionsCypherFactory) updateAttributeCypher(req acl.UpdateAttributeReq) (string, map[string]interface{}) {
	return ncUpdateAttributeQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"attrName":    req.Attribute.Name(),
			"attrKind":    req.Attribute.Kind(),
			"attrValue":   req.Attribute.Value(),
			"composition": model.CompositionRelationship}
}

const ncDeleteAttributeQuery = `
MATCH ((:Resource{name: $name})-[:Includes{kind: $composition}]->(a:Attribute{name: $attrName}))
DETACH DELETE a
`

func (f nonCachedPermissionsCypherFactory) deleteAttributeCypher(req acl.DeleteAttributeReq) (string, map[string]interface{}) {
	return ncDeleteAttributeQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"attrName":    req.AttributeId,
			"composition": model.CompositionRelationship}
}

const ncCreateRelQuery = `
MATCH (parent:Resource{name: $parentName})
MATCH (child:Resource{name: $childName})
WHERE NOT (parent)-[:Includes]->(child) AND NOT (child)-[:Includes*]->(parent)
CREATE (parent)-[:Includes{kind: $relKind}]->(child)
`

func (f nonCachedPermissionsCypherFactory) createAggregationRelCypher(req acl.CreateAggregationRelReq) (string, map[string]interface{}) {
	return ncCreateRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.AggregateRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

const ncDeleteRelQuery = `
MATCH (parent:Resource{name: $parentName})-[includes:Includes{kind: $relKind}]->(child:Resource{name: $childName})
DELETE includes
`

func (f nonCachedPermissionsCypherFactory) deleteAggregationRelCypher(req acl.DeleteAggregationRelReq) (string, map[string]interface{}) {
	return ncDeleteRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.AggregateRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

func (f nonCachedPermissionsCypherFactory) createCompositionRelCypher(req acl.CreateCompositionRelReq) (string, map[string]interface{}) {
	return ncCreateRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.CompositionRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

func (f nonCachedPermissionsCypherFactory) deleteCompositionRelCypher(req acl.DeleteCompositionRelReq) (string, map[string]interface{}) {
	return ncDeleteRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.CompositionRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

const ncCreatePermissionQuery = `
MATCH (sub:Resource{name: $subName})
MATCH (obj:Resource{name: $objName})
WHERE NOT ((sub)-[:Has]->(:Permission{name: $permName, kind: $permKind})-[:On]->(obj))
CREATE (sub)-[:Has]->(:Permission{name: $permName, kind: $permKind, condition: $permCond})-[:On]->(obj)
`

func (f nonCachedPermissionsCypherFactory) createPermissionCypher(req acl.CreatePermissionReq) (string, map[string]interface{}) {
	return ncCreatePermissionQuery,
		map[string]interface{}{
			"subName":  req.Subject.Name(),
			"objName":  req.Object.Name(),
			"permName": req.Permission.Name(),
			"permKind": req.Permission.Kind(),
			"permCond": req.Permission.Condition().Expression()}
}

const ncDeletePermissionQuery = `
MATCH (sub:Resource{name: $subName})
MATCH (obj:Resource{name: $objName})
MATCH ((sub)-[:Has]->(p:Permission{name: $permName, kind: $permKind})-[:On]->(obj))
DETACH DELETE p
`

func (f nonCachedPermissionsCypherFactory) deletePermissionCypher(req acl.DeletePermissionReq) (string, map[string]interface{}) {
	return ncDeletePermissionQuery,
		map[string]interface{}{
			"subName":  req.Subject.Name(),
			"objName":  req.Object.Name(),
			"permName": req.Permission.Name(),
			"permKind": req.Permission.Kind()}
}

const ncGetPermissionsQuery = `
OPTIONAL MATCH (sub:Resource{name: $subName})<-[:Includes*0..]-(subParent:Resource)-[:Has]->
(p:Permission{name: $permName})-[:On]->(objParent:Resource)-[:Includes*0..]->(obj:Resource{name: $objName})
WHERE sub IS NOT null AND obj IS NOT null AND p IS NOT null
WITH p, sub, subParent, obj, objParent
CALL {
	WITH sub, subParent
	MATCH path=(sub)<-[:Includes*0..100]-(subParent)
	RETURN -length(path) AS subPriority
	ORDER BY subPriority ASC
	LIMIT 1
}
CALL {
	WITH obj, objParent
	MATCH path=(obj)<-[:Includes*0..100]-(objParent)
	RETURN -length(path) AS objPriority
	ORDER BY objPriority ASC
	LIMIT 1
}
RETURN p.name, p.kind, p.condition, subPriority, objPriority
`

func (f nonCachedPermissionsCypherFactory) getEffectivePermissionsWithPriorityCypher(req acl.GetPermissionHierarchyReq) (string, map[string]interface{}) {
	return ncGetPermissionsQuery,
		map[string]interface{}{
			"subName":  req.Subject.Name(),
			"objName":  req.Object.Name(),
			"permName": req.PermissionName}
}

type cachedPermissionsCypherFactory struct {
}

func NewCachedPermissionsCypherFactory() cypherFactory {
	return &cachedPermissionsCypherFactory{}
}

const cCreateResourceQuery = `
MERGE (r:Resource{name: $name})
MERGE (root:Resource{name: $rootName})
MERGE (root)-[:Includes{kind: $composition}]->(r)
WITH r, root
CALL {
	WITH r, root
	MATCH (root)-[srel:Has]->(p:Permission)
	MERGE (r)-[:Has{priority: srel.prioriy - 1}]->(p)
}
CALL {
	WITH r, root
	MATCH (root)<-[orel:On]-(p:Permission)
	MERGE (r)<-[:On{priority: orel.priority - 1}]-(p)
}
`

func (f cachedPermissionsCypherFactory) createResourceCypher(req acl.CreateResourceReq) (string, map[string]interface{}) {
	return cCreateResourceQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"rootName":    model.RootResource.Name(),
			"composition": model.CompositionRelationship}
}

const cDeleteResourceCypher = `
MATCH (r:Resource{name: $name})
WITH r
// nadji sve resurse za brisanje
CALL {
    WITH r
    OPTIONAL MATCH (r)-[:Includes*{kind: $composition}]->(d:Resource)
    RETURN collect(d) + collect(r) AS delRes
}
// obrisi sve atribute resursa za brisanje
CALL {
    WITH delRes
    UNWIND delRes AS r
    MATCH (r)-[:Includes{kind: $composition}]->(a:Attribute)
    DETACH DELETE a
}
// obrisi sve direktno dodeljene dozvole resursa za brisanje
CALL {
    WITH delRes
    UNWIND delRes AS r
    MATCH (r)-[:Has|On{priority: 0}]-(p:Permission{})
    DETACH DELETE p
}
// nadji sve resurse kojima treba ukloniti dozvole
CALL {
    WITH delRes
    MATCH (p:Resource)-[:Includes*]->(c:Resource)
    WHERE p IN delRes AND NOT c IN delRes
    RETURN collect(p) AS permRes
}
// ukloni nasledjene dozvole iz potomaka
CALL {
    WITH permRes, delRes
    UNWIND permRes AS res
    // nadji i obrisi sve putanje resursa do direktno dodeljene dozvole - subjekat
    CALL {
      WITH res, delRes
      MATCH path=(res)<-[:Includes*]-(parent:Resource)-[:Has{priority: 0}]->(perm:Permission)
      WHERE ANY(pathRes IN NODES(path) WHERE pathRes IN delRes)
      MATCH ((res)-[srel:Has{priority: -(length(path)-1)}]->(perm))
      WITH collect(srel) AS del
      UNWIND del[..1] AS d
      DELETE d
    }
    // nadji i obrisi sve putanje resursa do direktno dodeljene dozvole - objekat
    CALL {
      WITH res, delRes
      MATCH path=(res)<-[:Includes*]-(parent:Resource)<-[:On{priority: 0}]-(perm:Permission)
      WHERE ANY(pathRes IN NODES(path) WHERE pathRes IN delRes)
      MATCH ((res)<-[orel:On{priority: -(length(path)-1)}]-(perm))
      WITH collect(orel) AS del
      UNWIND del[..1] AS d
      DELETE d
    }
}
// obrisi resurse
CALL {
    WITH delRes
    UNWIND delRes AS r
    DETACH DELETE r
}
`

func (f cachedPermissionsCypherFactory) deleteResourceCypher(req acl.DeleteResourceReq) (string, map[string]interface{}) {
	return cDeleteResourceCypher,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

const cGetResourceQuery = `
MATCH (resource:Resource{name: $name})
OPTIONAL MATCH (attr:Attribute)<-[:Includes{kind: $composition}]-(resource)
RETURN resource.name, collect(properties(attr)) as attrs
`

func (f cachedPermissionsCypherFactory) getResourceCypher(req acl.GetResourceReq) (string, map[string]interface{}) {
	return cGetResourceQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"composition": model.CompositionRelationship}
}

const cCreateAttributeQuery = `
MATCH (r:Resource{name: $name})
WHERE NOT ((r)-[:Includes{kind: $composition}]->(:Attribute{name: $attrName}))
CREATE (r)-[:Includes{kind: $composition}]->(:Attribute{name: $attrName, kind: $attrKind, value: $attrValue})
`

func (f cachedPermissionsCypherFactory) createAttributeCypher(req acl.CreateAttributeReq) (string, map[string]interface{}) {
	return cCreateAttributeQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"attrName":    req.Attribute.Name(),
			"attrKind":    req.Attribute.Kind(),
			"attrValue":   req.Attribute.Value(),
			"composition": model.CompositionRelationship}
}

const cUpdateAttributeQuery = `
MATCH ((:Resource{name: $name})-[:Includes{kind: $composition}]->(a:Attribute{name: $attrName, kind: $attrKind}))
SET a.value = $attrValue
`

func (f cachedPermissionsCypherFactory) updateAttributeCypher(req acl.UpdateAttributeReq) (string, map[string]interface{}) {
	return cUpdateAttributeQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"attrName":    req.Attribute.Name(),
			"attrKind":    req.Attribute.Kind(),
			"attrValue":   req.Attribute.Value(),
			"composition": model.CompositionRelationship}
}

const cDeleteAttributeQuery = `
MATCH ((:Resource{name: $name})-[:Includes{kind: $composition}]->(a:Attribute{name: $attrName}))
DETACH DELETE a
`

func (f cachedPermissionsCypherFactory) deleteAttributeCypher(req acl.DeleteAttributeReq) (string, map[string]interface{}) {
	return cDeleteAttributeQuery,
		map[string]interface{}{
			"name":        req.Resource.Name(),
			"attrName":    req.AttributeId,
			"composition": model.CompositionRelationship}
}

const cCreateRelQuery = `
MATCH (parent:Resource{name: $parentName})
MATCH (child:Resource{name: $childName})
WHERE NOT (parent)-[:Includes]->(child) AND NOT (child)-[:Includes*]->(parent)
// kreiraj novy vezu
CREATE (parent)-[newRel:Includes{kind: $relKind}]->(child)
// nadji sve dozvole koje treba da se naslede
WITH parent, newRel
CALL {
    WITH parent
    MATCH (parent)-[srel:Has|On]-(p:Permission)
    RETURN collect({priority: srel.priority, type: type(srel), permission: p}) AS rels
}
// nadji nove putanje od roditelja do potomaka
CALL {
    WITH parent, newRel
    MATCH path=(parent)-[:Includes*]->(d:Resource)
    WHERE newRel in RELATIONSHIPS(path)
    RETURN collect(path) AS newPaths
}
//  dodeli dozvole potomcima na novim putanjama
CALL {
    WITH newPaths, rels
    UNWIND newPaths AS newPath
    WITH last(nodes(newPath)) AS res, length(newPath) AS dist, rels AS policies
    UNWIND policies AS policy
    WITH policy.priority AS priority, policy.type AS type, policy.permission AS policy, res, dist
    FOREACH (i in CASE WHEN type = "Has" THEN [1] ELSE [] END |
        CREATE (res)-[:Has{priority: priority - dist}]->(policy)
    )
    FOREACH (i in CASE WHEN type = "On" THEN [1] ELSE [] END |
        CREATE (res)<-[:On{priority: priority - dist}]-(policy)
    )
}
`

func (f cachedPermissionsCypherFactory) createAggregationRelCypher(req acl.CreateAggregationRelReq) (string, map[string]interface{}) {
	return cCreateRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.AggregateRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

const cDeleteRelQuery = `
MATCH (parent:Resource{name: $parentName})-[includes:Includes{kind: $relKind}]->(child:Resource{name: $childName})
WITH parent, includes
// nadji sve dozvole koje ima roditelj (kao subjekat)
CALL {
    WITH parent
    MATCH (parent)-[rel:Has]->(p:Permission)
    RETURN collect({permission: p, priority: rel.priority}) AS subPermissions
}
// nadji sve dozvole koje ima roditelj (kao objekat)
CALL {
    WITH parent
    MATCH (parent)<-[rel:On]-(p:Permission)
    RETURN collect({permission: p, priority: rel.priority}) AS objPermissions
}
// pronadji sve putanje od roditelja koje ce biti prekinute
CALL {
    WITH parent, includes
    MATCH path=(parent)-[:Includes*]->(d:Resource)
    WHERE includes IN RELATIONSHIPS(path)
    RETURN collect(path) AS oldPaths
}
// obrisi sve nasledjene dozvole
CALL {
    WITH parent, subPermissions, objPermissions, oldPaths
    CALL {
        WITH parent, subPermissions, oldPaths
        UNWIND oldPaths AS oldPath
        WITH last(nodes(oldPath)) AS res, length(oldPath) AS dist, subPermissions
        UNWIND subPermissions AS policy
        WITH policy.priority AS priority, policy.permission AS permission, res, dist
        MATCH (res)-[rrel:Has{priority: priority - dist}]->(permission)
        WITH permission, rrel.priority as priority, collect(rrel) AS del
        UNWIND del[..1] AS d
        DELETE d
    }
    CALL {
        WITH parent, objPermissions, oldPaths
        UNWIND oldPaths AS oldPath
        WITH last(nodes(oldPath)) AS res, length(oldPath) AS dist, objPermissions
        UNWIND objPermissions AS policy
        WITH policy.priority AS priority, policy.permission AS permission, res, dist
        MATCH (res)<-[rrel:On{priority: priority - dist}]-(permission)
        WITH permission, rrel.priority as priority, collect(rrel) AS del
        UNWIND del[..1] AS d
        DELETE d
    }
}
// obrisi vezu izmedju roditelja i deteta
CALL {
    WITH includes
    DELETE includes
}
`

func (f cachedPermissionsCypherFactory) deleteAggregationRelCypher(req acl.DeleteAggregationRelReq) (string, map[string]interface{}) {
	return cDeleteRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.AggregateRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

func (f cachedPermissionsCypherFactory) createCompositionRelCypher(req acl.CreateCompositionRelReq) (string, map[string]interface{}) {
	return cCreateRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.CompositionRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

func (f cachedPermissionsCypherFactory) deleteCompositionRelCypher(req acl.DeleteCompositionRelReq) (string, map[string]interface{}) {
	return cDeleteRelQuery,
		map[string]interface{}{
			"parentName":  req.Parent.Name(),
			"childName":   req.Child.Name(),
			"relKind":     model.CompositionRelationship,
			"composition": model.CompositionRelationship,
			"rootName":    model.RootResource.Name()}
}

const cCreatePermissionQuery = `
MATCH (sub:Resource{name: $subName})
MATCH (obj:Resource{name: $objName})
WHERE NOT (sub)-[:Has{priority: 0}]->(:Permission{name: $permName, kind: $permKind})-[:On{priority: 0}]->(obj)
CREATE (sub)-[srel:Has{priority: 0}]->(p:Permission{name: $permName, kind: $permKind, condition: $permCond})-[orel:On{priority: 0}]->(obj)
WITH sub, obj, p
CALL {
    WITH sub, p
    MATCH path=((sub)-[:Includes*]->(descendant:Resource))
    CREATE (descendant)-[:Has{priority: -length(path)}]->(p)
}
CALL {
    WITH obj, p
    MATCH path=((obj)-[:Includes*]->(descendant:Resource))
    CREATE (descendant)<-[:On{priority: -length(path)}]-(p)
}
`

func (f cachedPermissionsCypherFactory) createPermissionCypher(req acl.CreatePermissionReq) (string, map[string]interface{}) {
	return cCreatePermissionQuery,
		map[string]interface{}{
			"subName":  req.Subject.Name(),
			"objName":  req.Object.Name(),
			"permName": req.Permission.Name(),
			"permKind": req.Permission.Kind(),
			"permCond": req.Permission.Condition().Expression()}
}

const cDeletePermissionQuery = `
MATCH (sub:Resource{name: $subName})
MATCH (obj:Resource{name: $objName})
MATCH (sub)-[:Has{priority: 0}]->(p:Permission{name: $permName, kind: $permKind})-[:On{priority: 0}]->(obj)
DETACH DELETE p
`

func (f cachedPermissionsCypherFactory) deletePermissionCypher(req acl.DeletePermissionReq) (string, map[string]interface{}) {
	return cDeletePermissionQuery,
		map[string]interface{}{
			"subName":  req.Subject.Name(),
			"objName":  req.Object.Name(),
			"permName": req.Permission.Name(),
			"permKind": req.Permission.Kind()}
}

const cGetPermissionsQuery = `
MATCH (sub:Resource{name: $subName})
MATCH (obj:Resource{name: $objName})
MATCH (sub)-[srel:Has]->(p:Permission{name: $permName})-[orel:On]->(obj)
WITH p, min(srel.priority) AS spriority, min(orel.priority) AS opriority
RETURN p.name, p.kind, p.condition, spriority, opriority
`

func (f cachedPermissionsCypherFactory) getEffectivePermissionsWithPriorityCypher(req acl.GetPermissionHierarchyReq) (string, map[string]interface{}) {
	return cGetPermissionsQuery,
		map[string]interface{}{
			"subName":  req.Subject.Name(),
			"objName":  req.Object.Name(),
			"permName": req.PermissionName}
}